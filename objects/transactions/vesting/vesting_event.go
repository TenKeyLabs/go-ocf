package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

// Object describing the transaction of an non-schedule-driven vesting event associated with
// a security
type VestingEvent struct {
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_EVENT`.
	VestingConditionID *string `json:"vesting_condition_id,omitempty"`
}
