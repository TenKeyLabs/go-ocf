package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/vesting"
)

// Describes a vesting condition satisfied when a period of time, relative to another
// vesting condition, has elapsed.
type VestingScheduleRelativeTrigger struct {
	vesting.VestingConditionTrigger

	// The span of time that must have elapsed since the condition `relative_to_condition_id`
	// occurred for this condition to trigger. For weeks or "ideal" years (365 days), use
	// `VestingPeriodInDays`. For calendar years use `VestingPeriodInMonths`.
	Period VestingPeriod `json:"period"`
	// Reference to the vesting condition ID to which the `period` is relative
	RelativeToConditionID string `json:"relative_to_condition_id"`
}
