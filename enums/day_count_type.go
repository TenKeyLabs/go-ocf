package enums

// How many days are there is a given period for calculation purposes?
//
// Enumeration of how the number of days are determined per period
type DayCountType string

const (
	Actual365 DayCountType = "ACTUAL_365"
	The30_360 DayCountType = "30_360"
)
