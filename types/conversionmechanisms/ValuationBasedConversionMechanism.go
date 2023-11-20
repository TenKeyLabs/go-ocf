package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism based on valuations
//
// Abstract type setting forth required field(s) for ALL conversion mechanism types
type ValuationBasedConversionMechanism struct {
	// How is company capitalization defined for purposes of exercise calculations? If possible,
	// include the legal language from the instrument.
	CapitalizationDefinition *string `json:"capitalization_definition,omitempty"`
	// The rules for which types of securities would be included in the capitalization
	// definition.
	CapitalizationDefinitionRules *types.CapitalizationDefinitionRules `json:"capitalization_definition_rules,omitempty"`
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
	// If there is a specified valuation figure to use, what is it? Look to `valuation_type` to
	// understand whether this represents, a max valuation (`CAP`), actual valuation at time of
	// exercise (`ACTUAL`) or fixed valuation (`FIXED`).
	ValuationAmount *types.Monetary                 `json:"valuation_amount,omitempty"`
	ValuationType   enums.ValuationBasedFormulaType `json:"valuation_type"`
}
