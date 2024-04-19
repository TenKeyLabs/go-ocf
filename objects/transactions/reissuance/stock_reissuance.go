package reissuance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/reissuance"
)

// Object describing a re-issuance of stock
type StockReissuance struct {
	reissuance.Reissuance
	transactions.SecurityTransaction
}
