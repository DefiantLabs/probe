package query

import (
	protorevTypes "github.com/nodersteam/probe/client/codec/osmosis/v15/x/protorev/types"
	queryTypes "github.com/nodersteam/probe/query"
)

func ProtorevDeveloperAccountRPC(q *queryTypes.Query) (*protorevTypes.QueryGetProtoRevDeveloperAccountResponse, error) {
	req := protorevTypes.QueryGetProtoRevDeveloperAccountRequest{}
	queryClient := protorevTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.GetProtoRevDeveloperAccount(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
