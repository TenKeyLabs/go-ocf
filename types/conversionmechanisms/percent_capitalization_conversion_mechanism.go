package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism of percent of capitalization conversion (where
// an instrument purports to grant a percent of company capitalization at some point in
// time)
type PercentCapitalizationConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// How is company capitalization defined for purposes of conversion? If possible, include
	// the legal language from the instrument.
	CapitalizationDefinition *string `json:"capitalization_definition,omitempty"`
	// The rules for which types of securities would be included in the capitalization
	// definition.
	CapitalizationDefinitionRules *types.CapitalizationDefinitionRules `json:"capitalization_definition_rules,omitempty"`
	// What percentage of the company capitalization does this convert to
	ConvertsToPercent string `json:"converts_to_percent"`
}
