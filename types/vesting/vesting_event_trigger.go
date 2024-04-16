package vesting

import "github.com/tenkeylabs/go-ocf/primitives/types/vesting"

// Describes a vesting condition satisfied when a particular unscheduled event occurs
type VestingEventTrigger struct {
	vesting.VestingConditionTrigger
}
