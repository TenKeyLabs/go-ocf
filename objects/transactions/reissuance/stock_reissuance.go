package reissuance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/reissuance"
)

// Object describing a re-issuance of stock
type StockReissuance struct {
	reissuance.Reissuance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}
