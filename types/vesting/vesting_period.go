package vesting

import (
	"encoding/json"
	"errors"

	"github.com/tenkeylabs/go-ocf/enums"
)

type VestingPeriod struct {
	VestingPeriodInDays   *VestingPeriodInDays   `json:"vesting_period_in_days,omitempty"`
	VestingPeriodInMonths *VestingPeriodInMonths `json:"vesting_period_in_months,omitempty"`
}

func (v *VestingPeriod) UnmarshalJSON(data []byte) error {
	aliasVestingPeriod := &struct {
		DayOfMonth *enums.VestingDayOfMonth `json:"day_of_month"`
	}{}
	if err := json.Unmarshal(data, &aliasVestingPeriod); err != nil {
		return err
	}

	var err error
	if aliasVestingPeriod.DayOfMonth == nil {
		var vestingPeriodInDays VestingPeriodInDays
		err = json.Unmarshal(data, &vestingPeriodInDays)
		v.VestingPeriodInDays = &vestingPeriodInDays
	} else {
		var vestingPeriodInMonths VestingPeriodInMonths
		err = json.Unmarshal(data, &vestingPeriodInMonths)
		v.VestingPeriodInMonths = &vestingPeriodInMonths
	}

	return err
}

func (v VestingPeriod) MarshalJSON() ([]byte, error) {
	if v.VestingPeriodInDays != nil {
		return json.Marshal(v.VestingPeriodInDays)
	} else if v.VestingPeriodInMonths != nil {
		return json.Marshal(v.VestingPeriodInMonths)
	}

	return nil, errors.New("unknown vesting period")
}
