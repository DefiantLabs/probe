// APACHE NOTICE
// Source copied and modified from https://github.com/strangelove-ventures/lens
package query

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

type QueryOptions struct {
	Pagination *query.PageRequest
	Height     int64
}
