// APACHE NOTICE
// Source copied and modified from https://github.com/strangelove-ventures/lens
package query

import (
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func BlockResultsRPC(q *Query) (*coretypes.ResultBlockResults, error) {
	var height int64
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	// If height is not specified, default value is 0, query the latest available block then
	if q.Options.Height == 0 {
		resStatus, err := q.Client.RPCClient.Status(ctx)
		if err != nil {
			return nil, err
		}
		height = resStatus.SyncInfo.LatestBlockHeight
	} else {
		height = q.Options.Height
	}
	res, err := q.Client.RPCClient.BlockResults(ctx, &height)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func BlockRPC(q *Query) (*coretypes.ResultBlock, error) {
	var height int64
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	// If height is not specified, default value is 0, query the latest available block then
	if q.Options.Height == 0 {
		resStatus, err := q.Client.RPCClient.Status(ctx)
		if err != nil {
			return nil, err
		}
		height = resStatus.SyncInfo.LatestBlockHeight
	} else {
		height = q.Options.Height
	}
	res, err := q.Client.RPCClient.Block(ctx, &height)
	if err != nil {
		return nil, err
	}

	return res, nil
}
