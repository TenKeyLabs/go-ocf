package vesting

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/types/vesting"
)

// Describes a period of time expressed in months (e.g. 3 months) for use in Vesting Terms.
type VestingPeriodInMonths struct {
	vesting.VestingPeriod

	// The calendar day of a month to award vesting.
	DayOfMonth enums.VestingDayOfMonth `json:"day_of_month"`
}
