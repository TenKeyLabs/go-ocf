package acceptance

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/acceptance"
)

// Object describing equity compensation acceptance transaction
type EquityCompensationAcceptance struct {
	transactions.SecurityTransaction

	Acceptance acceptance.Acceptance `json:"acceptance,omitempty"`
}
