package staking

import (
	probeQueryTypes "github.com/DefiantLabs/probe/query"
	queryTypes "github.com/cosmos/cosmos-sdk/types/query"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func DelegatorDelegationsRPC(q *probeQueryTypes.Query, delegatorAddress string, paginationKey []byte) (*stakingTypes.QueryDelegatorDelegationsResponse, error) {
	req := stakingTypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegatorAddress,
	}

	if paginationKey != nil {
		req.Pagination = &queryTypes.PageRequest{
			Key: paginationKey,
		}
	}

	queryClient := stakingTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DelegatorDelegations(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DelegatorDelegations(q *probeQueryTypes.Query, address string) (*stakingTypes.QueryDelegatorDelegationsResponse, error) {
	var paginationKey []byte

	if q.Options.Pagination != nil && q.Options.Pagination.Key != nil {
		paginationKey = q.Options.Pagination.Key
	}

	return DelegatorDelegationsRPC(q, address, paginationKey)
}
