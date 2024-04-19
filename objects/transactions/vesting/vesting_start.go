package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

// Object describing the transaction of vesting schedule start / commencement associated
// with a security
type VestingStart struct {
	transactions.SecurityTransaction

	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_EVENT`.
	VestingConditionID *string `json:"vesting_condition_id,omitempty"`
}
