package enums

// What is the period over which interest is calculated?
//
// Enumeration of interest accrual period types
type AccrualPeriodType string

const (
	Annual     AccrualPeriodType = "ANNUAL"
	Daily      AccrualPeriodType = "DAILY"
	Monthly    AccrualPeriodType = "MONTHLY"
	Quarterly  AccrualPeriodType = "QUARTERLY"
	SemiAnnual AccrualPeriodType = "SEMI_ANNUAL"
)
