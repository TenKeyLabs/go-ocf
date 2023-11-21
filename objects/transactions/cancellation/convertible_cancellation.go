package cancellation

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/cancellation"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a cancellation of a convertible security
type ConvertibleCancellation struct {
	cancellation.Cancellation
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Amount of monetary value cancelled
	Amount types.Monetary `json:"amount"`
}
