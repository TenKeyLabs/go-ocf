package vesting

import "github.com/tenkeylabs/go-ocf/enums"

// Describes a period of time expressed in months (e.g. 3 months) for use in Vesting Terms.
//
// Abstract type describing the fields common to all periods of time (e.g. 3 months, 365
// days) for use in Vesting Terms
type VestingPeriodInMonths struct {
	// The calendar day of a month to award vesting.
	DayOfMonth enums.VestingDayOfMonth `json:"day_of_month"`
	// The quantity of `type` units of time; e.g. for 3 months, this would be `3`; for 30 days,
	// this would be `30`
	Length int64 `json:"length"`
	// The number of times this vesting period triggers. If vesting occurs monthly for 36
	// months, for example, this would be `36`
	Occurrences int64 `json:"occurrences"`
	// The unit of time for the period, e.g. `MONTHS` or `DAYS`
	Type enums.PeriodType `json:"type"`
}
