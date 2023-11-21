package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a warrant acceptance transaction
type WarrantAcceptance struct {
	acceptance.Acceptance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}
