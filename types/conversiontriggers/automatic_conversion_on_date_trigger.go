package conversiontriggers

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/conversiontriggers"
)

// Type representation of an automatic trigger on a date.
type AutomaticConversionOnDateTrigger struct {
	conversiontriggers.ConversionTrigger

	// Date on which trigger occurs automatically (if it hasn't already occured)
	TriggerDate string `json:"trigger_date"`
}
