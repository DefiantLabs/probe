package main

import (
	"fmt"
	"os"

	"github.com/DefiantLabs/probe/client"
	gammTypes "github.com/DefiantLabs/probe/client/codec/osmosis/v15/x/gamm/types"
	querier "github.com/DefiantLabs/probe/query"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	cquery "github.com/cosmos/cosmos-sdk/types/query"
	ibcChanTypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

func main() {

	cconfig := &client.ChainClientConfig{
		Key:            "default",
		ChainID:        os.Getenv("CHAIN_ID"),
		RPCAddr:        os.Getenv("RPC_SERVER"),
		AccountPrefix:  os.Getenv("ACCOUNT_PREFIX"),
		KeyringBackend: "test",
		Debug:          false,
		Timeout:        "60s",
		OutputFormat:   "json",
		Modules:        client.ModuleBasics,
	}
	cl, err := client.NewChainClient(cconfig, "", nil, nil)
	if err != nil {
		fmt.Println("Error setting up client")
		fmt.Println(err)
		os.Exit(1)
	}

	// Proof of concept code
	var checkHeight int64 = 9736414

	// Check chain status
	query := querier.Query{Client: cl, Options: &querier.QueryOptions{}}
	status, err := querier.StatusRPC(&query)
	if err != nil {
		fmt.Println("Error getting status")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Got status, some data follows:")
		fmt.Printf("Node moniker: %+v\n", status.NodeInfo.Moniker)
	}

	// Get a block
	options := querier.QueryOptions{Height: checkHeight}
	query = querier.Query{Client: cl, Options: &querier.QueryOptions{Height: checkHeight}}
	block, err := querier.BlockRPC(&query)
	if err != nil {
		fmt.Println("Error getting block")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Got block, some data follows:")
		fmt.Printf("Height: %+v\n", block.Block.Height)
	}

	// Get block results
	options = querier.QueryOptions{Height: checkHeight}
	query = querier.Query{Client: cl, Options: &querier.QueryOptions{Height: checkHeight}}
	blockResults, err := querier.BlockResultsRPC(&query)
	if err != nil {
		fmt.Println("Error getting block results")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Got block results, some data follows:")
		fmt.Printf("Height: %+v\n", blockResults.Height)
	}

	// Get some Transactions at a specific height
	pg := cquery.PageRequest{Limit: 100}
	options = querier.QueryOptions{Height: checkHeight, Pagination: &pg}
	query = querier.Query{Client: cl, Options: &options}

	txResponse, err := querier.TxsAtHeightRPC(&query, checkHeight, cl.Codec)

	if err != nil {
		fmt.Println("Error getting transactions")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Got txes, some TX hashes follow:")
		for i := range txResponse.Txs {
			currTx := txResponse.Txs[i]
			currTxResp := txResponse.TxResponses[i]
			fmt.Println("TX Hash:", currTxResp.TxHash)
			fmt.Println("Contains these messages:")
			for msgIdx := range currTx.Body.Messages {
				currMsg := currTx.Body.Messages[msgIdx].GetCachedValue()
				if currMsg == nil {
					// This happens if the Codec for the client is missing a message type
					fmt.Println("Error getting CachedValue for", currTx.Body.Messages[msgIdx].TypeUrl)
				} else {
					// Pass the message to some handler function
					handler, ok := handlers[currTx.Body.Messages[msgIdx].TypeUrl]
					if ok {
						fmt.Println("Found handler for ", currTx.Body.Messages[msgIdx].TypeUrl)
						handler(currMsg.(cosmosTypes.Msg))
					} else {
						fmt.Println("No handler for ", currTx.Body.Messages[msgIdx].TypeUrl)
					}
				}
			}
		}
	}
}

var handlers = map[string]func(cosmosTypes.Msg){
	"/osmosis.gamm.v1beta1.MsgSwapExactAmountOut": func(currMsg cosmosTypes.Msg) {
		swapExactAmountOut := currMsg.(*gammTypes.MsgSwapExactAmountOut)
		fmt.Printf("%s swapped %s\n", swapExactAmountOut.Sender, swapExactAmountOut.TokenOut)
	},
	"/ibc.core.channel.v1.MsgAcknowledgement": func(currMsg cosmosTypes.Msg) {
		ack := currMsg.(*ibcChanTypes.MsgAcknowledgement)
		fmt.Printf("%s acked with result %s\n", ack.Signer, ack.Acknowledgement)
	},
}
