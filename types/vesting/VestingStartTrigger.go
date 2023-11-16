package vesting

import "github.com/tenkeylabs/go-ocf/enums"

// Describes a vesting condition satisfied at the security's vesting commencement date
//
// Abstract type describing fields needed in all triggers types, with a 'trigger' being a
// condition that must be satisfied for a VestingCondition to be met
type VestingStartTrigger struct {
	// Identifies the sub-type of trigger
	Type enums.VestingTriggerType `json:"type"`
}
