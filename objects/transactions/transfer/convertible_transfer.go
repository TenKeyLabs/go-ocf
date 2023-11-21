package transfer

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/transfer"
)

// Object describing the transaction of vesting schedule start / commencement associated
// with a security
type ConvertibleTransfer struct {
	transfer.Transfer
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_START_DATE`.
	VestingConditionID string `json:"vesting_condition_id"`
}
