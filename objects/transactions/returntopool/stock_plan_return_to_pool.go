package returntopool

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/returntopool"
)

// Object describing which stock plan pool a particular security's shares were returned to
// upon cancellation.
type StockPlanReturnToPool struct {
	returntopool.ReturnToPool
	transactions.SecurityTransaction
}
