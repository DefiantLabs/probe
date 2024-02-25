package query

import (
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func ValidatorsAtHeightRPC(q *Query, height int64) (*stakingTypes.QueryValidatorsResponse, error) {
	req := stakingTypes.QueryValidatorsRequest{}
	queryClient := stakingTypes.NewQueryClient(q.Client)

	if height > 0 {
		q.Options.Height = height
	}

	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Validators(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
