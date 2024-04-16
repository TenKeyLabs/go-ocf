package conversiontriggers

import "github.com/tenkeylabs/go-ocf/primitives/types/conversiontriggers"

// Type representation of elective trigger on fulfillment of a condition.
type ElectiveConversionOnConditionTrigger struct {
	conversiontriggers.ConversionTrigger

	// Legal language describing what conditions must be satisfied for the conversion to take
	// place (ideally, this should be excerpted from the instrument where possible)
	TriggerCondition string `json:"trigger_condition"`
}
