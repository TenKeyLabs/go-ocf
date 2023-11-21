package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a stock acceptance transaction
type StockAcceptance struct {
	acceptance.Acceptance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}
