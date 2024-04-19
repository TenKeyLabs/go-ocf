package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a stock acceptance transaction
type StockAcceptance struct {
	transactions.SecurityTransaction

	Acceptance acceptance.Acceptance `json:"acceptance,omitempty"`
}
