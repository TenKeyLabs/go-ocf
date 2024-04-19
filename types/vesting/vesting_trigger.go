package vesting

import (
	"encoding/json"
	"fmt"

	"github.com/tenkeylabs/go-ocf/enums"
)

type VestingTrigger struct {
	Type enums.VestingTriggerType `json:"type,omitempty"`

	VestingStartTrigger            *VestingStartTrigger
	VestingScheduleAbsoluteTrigger *VestingScheduleAbsoluteTrigger
	VestingScheduleRelativeTrigger *VestingScheduleRelativeTrigger
	VestingEventTrigger            *VestingEventTrigger
}

func (v *VestingTrigger) UnmarshalJSON(data []byte) error {
	aliasVestingTrigger := &struct {
		Type enums.VestingTriggerType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &aliasVestingTrigger); err != nil {
		return err
	}

	v.Type = aliasVestingTrigger.Type

	var err error
	switch aliasVestingTrigger.Type {
	case enums.VestingStartDate:
		var vestingStartTrigger VestingStartTrigger
		err = json.Unmarshal(data, &vestingStartTrigger)
		v.VestingStartTrigger = &vestingStartTrigger
	case enums.VestingScheduleAbsolute:
		var vestingScheduleAbsoluteTrigger VestingScheduleAbsoluteTrigger
		err = json.Unmarshal(data, &vestingScheduleAbsoluteTrigger)
		v.VestingScheduleAbsoluteTrigger = &vestingScheduleAbsoluteTrigger
	case enums.VestingScheduleRelative:
		var vestingScheduleRelativeTrigger VestingScheduleRelativeTrigger
		err = json.Unmarshal(data, &vestingScheduleRelativeTrigger)
		v.VestingScheduleRelativeTrigger = &vestingScheduleRelativeTrigger
	case enums.VestingEvent:
		var vestingEventTrigger VestingEventTrigger
		err = json.Unmarshal(data, &vestingEventTrigger)
		v.VestingEventTrigger = &vestingEventTrigger
	}

	return err
}

func (v VestingTrigger) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case enums.VestingStartDate:
		return json.Marshal(v.VestingStartTrigger)
	case enums.VestingScheduleAbsolute:
		return json.Marshal(v.VestingScheduleAbsoluteTrigger)
	case enums.VestingScheduleRelative:
		return json.Marshal(v.VestingScheduleRelativeTrigger)
	case enums.VestingEvent:
		return json.Marshal(v.VestingEventTrigger)
	}

	return nil, fmt.Errorf("unknown vesting trigger type: %s", v.Type)
}
