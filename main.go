package main

import (
	"fmt"
	"os"

	"github.com/DefiantLabs/probe/client"
	querier "github.com/DefiantLabs/probe/query"
	cquery "github.com/cosmos/cosmos-sdk/types/query"
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
		fmt.Println(err)
	}

	// Proof of concept code
	var checkHeight int64 = 15289176

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
					fmt.Println("Error getting CachedValue for", currTx.Body.Messages[msgIdx].TypeUrl)
				} else {
					fmt.Println(currTx.Body.Messages[msgIdx].TypeUrl)
				}
			}
		}
	}
}
