package vesting

import "github.com/tenkeylabs/go-ocf/primitives/types/vesting"

// Describes a vesting condition satisfied on an absolute date.
type VestingScheduleAbsoluteTrigger struct {
	vesting.VestingConditionTrigger

	//The date on which this condition triggers.
	Date string `json:"date"`
}
