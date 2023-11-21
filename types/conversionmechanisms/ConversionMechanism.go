package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

type ConversionMechanism struct {
	*CustomConversionMechanism
	*FixedAmountConversionMechanism
	*NoteConversionMechanism
	*PercentCapitalizationConversionMechanism
	*RatioConversionMechanism
	*SAFEConversionMechanism
	*SharePriceBasedConversionMechanism
	*ValuationBasedConversionMechanism
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"` //TODO: Remove type from each conversion mechanism
	// Is this an MFN (Most Favored Nations) flavored Convertible Note?
	ConversionMfn *bool `json:"conversion_mfn,omitempty"`
	// The rules for which types of securities would be included in the capitalization
	// definition.
	CapitalizationDefinitionRules *types.CapitalizationDefinitionRules `json:"capitalization_definition_rules,omitempty"`
}
