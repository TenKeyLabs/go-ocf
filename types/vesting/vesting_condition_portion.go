package vesting

// Describes a fractional portion (ratio) of shares associated with a Vesting Condition
type VestingConditionPortion struct {
	// Denominator of the ratio, i.e. the ratio of A to B (A:B) can be expressed as a fraction
	// (A/B), where B is the denominator
	Denominator string `json:"denominator"`
	// Numerator of the ratio, i.e. the ratio of A to B (A:B) can be expressed as a fraction
	// (A/B), where A is the numerator
	Numerator string `json:"numerator"`
	// If false, the ratio is applied to the entire quantity of the security's issuance. If
	// true, it is applied to the amount that has yet to vest. For example:
	// A stakeholder has been granted 1000 shares, and 400 are already vested.
	// If the portion is 1/5 and `remainder` is `false` for a VestingCondition, then that
	// condition will vest 200 shares -- 1/5 of the 1000 granted.
	// If the portion is 1/5 and `remainder` is `true`, then that condition will vest 120 shares
	// -- 1/5 of the 600 unvested.
	Remainder *bool `json:"remainder,omitempty"`
}
