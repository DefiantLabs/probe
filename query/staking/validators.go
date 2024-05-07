package staking

import (
	probeQueryTypes "github.com/DefiantLabs/probe/query"
	queryTypes "github.com/cosmos/cosmos-sdk/types/query"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func ValidatorsRPC(q *probeQueryTypes.Query, status stakingTypes.BondStatus, paginationKey []byte) (*stakingTypes.QueryValidatorsResponse, error) {
	req := stakingTypes.QueryValidatorsRequest{
		Status: status.String(),
	}

	if paginationKey != nil {
		req.Pagination = &queryTypes.PageRequest{
			Key: paginationKey,
		}
	}

	queryClient := stakingTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Validators(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Validators(q *probeQueryTypes.Query, status stakingTypes.BondStatus) (*stakingTypes.QueryValidatorsResponse, error) {
	var paginationKey []byte

	if q.Options.Pagination != nil && q.Options.Pagination.Key != nil {
		paginationKey = q.Options.Pagination.Key
	}

	return ValidatorsRPC(q, status, paginationKey)
}
