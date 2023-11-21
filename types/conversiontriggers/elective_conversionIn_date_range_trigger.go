package conversiontriggers

import "github.com/tenkeylabs/go-ocf/primitives/types/conversiontriggers"

// Type representation of elective trigger valid on or after start_date and until or before
// end_date.
type ElectiveConversionInDateRangeTrigger struct {
	conversiontriggers.ConversionTrigger

	// End date of range (inclusive)
	EndDate string `json:"end_date"`
	// Start date of range (inclusive)
	StartDate string `json:"start_date"`
}
