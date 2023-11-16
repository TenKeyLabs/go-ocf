package enums

// The type of period being measured (e.g. days or month)
//
// Enumeration of time period types
type PeriodType string

const (
	Days   PeriodType = "DAYS"
	Months PeriodType = "MONTHS"
	Years  PeriodType = "YEARS"
)
