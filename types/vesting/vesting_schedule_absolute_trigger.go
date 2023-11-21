package vesting

import "github.com/tenkeylabs/go-ocf/primitives/types/vesting"

// Describes a vesting condition satisfied on an absolute date.
type VestingScheduleAbsoluteTrigger struct {
	vesting.VestingConditionTrigger
}
