package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a convertible acceptance transaction
type ConvertibleAcceptance struct {
	acceptance.Acceptance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}
