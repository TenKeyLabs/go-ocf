package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a convertible acceptance transaction
type ConvertibleAcceptance struct {
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	Acceptance acceptance.Acceptance `json:"acceptance,omitempty"`
}
