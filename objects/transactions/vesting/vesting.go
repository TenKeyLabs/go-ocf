package vesting

type Vesting struct {
	*VestingAcceleration
	*VestingStart
	*VestingEvent

	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_EVENT`.
	VestingConditionID *string `json:"vesting_condition_id,omitempty"`
}
