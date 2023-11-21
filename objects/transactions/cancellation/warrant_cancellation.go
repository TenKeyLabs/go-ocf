package cancellation

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/cancellation"
)

// Object describing a cancellation of a warrant security
type WarrantCancellation struct {
	cancellation.Cancellation
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Quantity of non-monetary security units cancelled
	Quantity string `json:"quantity"`
}
