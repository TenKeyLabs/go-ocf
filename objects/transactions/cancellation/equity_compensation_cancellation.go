package cancellation

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/cancellation"
)

// Object describing a cancellation of equity compensation
type EquityCompensationCancellation struct {
	cancellation.Cancellation
	transactions.SecurityTransaction

	// Quantity of non-monetary security units cancelled
	Quantity string `json:"quantity"`
}
