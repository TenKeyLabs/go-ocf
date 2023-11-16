package vesting

// Describes condition / triggers to be satisfied for vesting to occur
type VestingCondition struct {
	// Detailed description of the condition
	Description *string `json:"description,omitempty"`
	// Reference identifier for this condition
	ID string `json:"id"`
	// List of ALL VestingCondition IDs that can trigger after this one. If there are none, use
	// an empty array.
	// Conditions should be in priority order in the array, ordered from the highest priority to
	// the lowest.
	NextConditionIDS []string `json:"next_condition_ids"`
	// If specified, the fractional part of the whole security that is vested, e.g. 25:100 for
	// 25%. Use `quantity` for a fixed vesting amount.
	Portion *VestingConditionPortion `json:"portion,omitempty"`
	// If specified, the fixed amount of the whole security to vest, e.g. 10000 shares. Use
	// `portion` for a proportional vesting amount.
	Quantity *string `json:"quantity,omitempty"`
	// Describes how this vesting condition is met, resulting in vesting the specified tranche
	// of shares
	Trigger VestingTrigger `json:"trigger"`
}
