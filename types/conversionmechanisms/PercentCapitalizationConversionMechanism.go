package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism of percent of capitalization conversion (where
// an instrument purports to grant a percent of company capitalization at some point in
// time)
//
// Abstract type setting forth required field(s) for ALL conversion mechanism types
type PercentCapitalizationConversionMechanism struct {
	// How is company capitalization defined for purposes of conversion? If possible, include
	// the legal language from the instrument.
	CapitalizationDefinition *string `json:"capitalization_definition,omitempty"`
	// The rules for which types of securities would be included in the capitalization
	// definition.
	CapitalizationDefinitionRules *types.CapitalizationDefinitionRules `json:"capitalization_definition_rules,omitempty"`
	// What percentage of the company capitalization does this convert to
	ConvertsToPercent string `json:"converts_to_percent"`
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
}
