package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing a warrant acceptance transaction
type WarrantAcceptance struct {
	transactions.SecurityTransaction

	Acceptance acceptance.Acceptance `json:"acceptance,omitempty"`
}
