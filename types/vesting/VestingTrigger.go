package vesting

import "github.com/tenkeylabs/go-ocf/enums"

type VestingTrigger struct {
	*VestingStartTrigger
	*VestingScheduleAbsoluteTrigger
	*VestingScheduleRelativeTrigger
	*VestingEventTrigger
	// Identifies the sub-type of trigger
	Type enums.VestingTriggerType `json:"type"`
}
