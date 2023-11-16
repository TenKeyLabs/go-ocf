package vesting

import "github.com/tenkeylabs/go-ocf/enums"

// Describes a vesting condition satisfied on an absolute date.
//
// Abstract type describing fields needed in all triggers types, with a 'trigger' being a
// condition that must be satisfied for a VestingCondition to be met
type VestingScheduleAbsoluteTrigger struct {
	// The date on which this condition triggers.
	Date string `json:"date"`
	// Identifies the sub-type of trigger
	Type enums.VestingTriggerType `json:"type"`
}
