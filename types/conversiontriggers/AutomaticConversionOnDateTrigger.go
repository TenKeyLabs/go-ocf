package conversiontriggers

import (
	"github.com/tenkeylabs/go-ocf/types/conversionrights"
	"github.com/tenkeylabs/go-ocf/types/vesting"
)

// Type representation of an automatic trigger on a date.
//
// Abstract type representation of required fields require for conversion rights types.
type AutomaticConversionOnDateTrigger struct {
	// When the conditions of the trigger are met, how does the convertible convert?
	ConversionRight conversionrights.ConversionRights `json:"conversion_right"`
	// Human-friendly nickname to describe the conversion right
	Nickname *string `json:"nickname,omitempty"`
	// Date on which trigger occurs automatically (if it hasn't already occured)
	TriggerDate string `json:"trigger_date"`
	// Long-form description of the trigger
	TriggerDescription *string `json:"trigger_description,omitempty"`
	// Id for this conversion trigger, unique within list of ConversionTriggers in parent
	// convertible issuance's `conversion_triggers` field.
	TriggerID string `json:"trigger_id"`
	// When the trigger condition is met, is the conversion automatic, elective or automatic
	// with an elective right not to convert
	Type vesting.VestingTrigger `json:"type"`
}
