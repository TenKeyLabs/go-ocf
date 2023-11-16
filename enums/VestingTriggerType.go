package enums

// Identifies the sub-type of trigger
//
// Enumeration of vesting trigger types
type VestingTriggerType string

const (
	VestingEvent            VestingTriggerType = "VESTING_EVENT"
	VestingScheduleAbsolute VestingTriggerType = "VESTING_SCHEDULE_ABSOLUTE"
	VestingScheduleRelative VestingTriggerType = "VESTING_SCHEDULE_RELATIVE"
	VestingStartDate        VestingTriggerType = "VESTING_START_DATE"
)
