// APACHE NOTICE
// Source copied and modified from https://github.com/strangelove-ventures/lens
package query

import (
	"context"
	"strconv"
	"time"

	"github.com/DefiantLabs/probe/client"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc/metadata"
)

type Query struct {
	Client  *client.ChainClient
	Options *QueryOptions
}

// GetQueryContext returns a context that includes the height and uses the timeout from the config
func (q *Query) GetQueryContext() (context.Context, context.CancelFunc) {
	timeout, _ := time.ParseDuration(q.Client.Config.Timeout) // Timeout is validated in the config so no error check
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	strHeight := strconv.Itoa(int(q.Options.Height))
	ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	return ctx, cancel
}

func (q *Query) BlockResults() (*coretypes.ResultBlockResults, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BlockResultsRPC(q)
}

func (q *Query) Block() (*coretypes.ResultBlock, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BlockRPC(q)
}

// Tx returns the Tx and all contained messages/TxResponse.
func (q *Query) TxByHeight(cc client.Codec) (*txTypes.GetTxsEventResponse, error) {
	return TxsAtHeightRPC(q, q.Options.Height, cc)
}

// Status returns information about a node status
func (q *Query) Status() (*coretypes.ResultStatus, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return StatusRPC(q)
}
