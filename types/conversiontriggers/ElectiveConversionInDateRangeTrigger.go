package conversiontriggers

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/conversionrights"
)

// Type representation of elective trigger valid on or after start_date and until or before
// end_date.
//
// Abstract type representation of required fields require for conversion rights types.
type ElectiveConversionInDateRangeTrigger struct {
	// When the conditions of the trigger are met, how does the convertible convert?
	ConversionRight conversionrights.ConversionRights `json:"conversion_right"`
	// End date of range (inclusive)
	EndDate string `json:"end_date"`
	// Human-friendly nickname to describe the conversion right
	Nickname *string `json:"nickname,omitempty"`
	// Start date of range (inclusive)
	StartDate string `json:"start_date"`
	// Long-form description of the trigger
	TriggerDescription *string `json:"trigger_description,omitempty"`
	// Id for this conversion trigger, unique within list of ConversionTriggers in parent
	// convertible issuance's `conversion_triggers` field.
	TriggerID string `json:"trigger_id"`
	// When the trigger condition is met, is the conversion automatic, elective or automatic
	// with an elective right not to convert
	Type enums.ConversionTriggerType `json:"type"`
}
