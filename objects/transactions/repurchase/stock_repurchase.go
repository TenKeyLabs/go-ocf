package repurchase

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/repurchase"
)

// Object describing a stock repurchase transaction
type StockRepurchase struct {
	repurchase.Repurchase
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}
