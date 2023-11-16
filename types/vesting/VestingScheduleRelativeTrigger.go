package vesting

import "github.com/tenkeylabs/go-ocf/enums"

// Describes a vesting condition satisfied when a period of time, relative to another
// vesting condition, has elapsed.
//
// Abstract type describing fields needed in all triggers types, with a 'trigger' being a
// condition that must be satisfied for a VestingCondition to be met
type VestingScheduleRelativeTrigger struct {
	// The span of time that must have elapsed since the condition `relative_to_condition_id`
	// occurred for this condition to trigger. For weeks or "ideal" years (365 days), use
	// `VestingPeriodInDays`. For calendar years use `VestingPeriodInMonths`.
	Period VestingPeriod `json:"period"`
	// Reference to the vesting condition ID to which the `period` is relative
	RelativeToConditionID string `json:"relative_to_condition_id"`
	// Identifies the sub-type of trigger
	Type enums.VestingTriggerType `json:"type"`
}
