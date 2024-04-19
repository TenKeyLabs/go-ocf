package transfer

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/transfer"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing the transaction of vesting schedule start / commencement associated
// with a security
type ConvertibleTransfer struct {
	transfer.Transfer
	transactions.SecurityTransaction

	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_START_DATE`.
	VestingConditionID string `json:"vesting_condition_id,omitempty"`
	//Amount of monetary value transferred
	Amount types.Monetary `json:"amount"`
}
