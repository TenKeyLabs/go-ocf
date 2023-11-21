package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism of a SAFE (mirrors the flavors and inputs of
// the Y Combinator SAFE)
//
// Abstract type setting forth required field(s) for ALL conversion mechanism types
type SAFEConversionMechanism struct {
	// How is company capitalization defined for purposes of conversion? If possible, include
	// the legal language from the instrument.
	CapitalizationDefinition *string `json:"capitalization_definition,omitempty"`
	// What is the percentage discount available upon conversion, if applicable? (decimal
	// representation - e.g. 0.125 for 12.5%)
	ConversionDiscount *string `json:"conversion_discount,omitempty"`
	// Is this an MFN flavored SAFE?
	ConversionMfn bool `json:"conversion_mfn"`
	// Should the conversion amount be based on pre or post money capitalization
	ConversionTiming *enums.ConversionTimingType `json:"conversion_timing,omitempty"`
	// What is the valuation cap (if applicable)?
	ConversionValuationCap *types.Monetary `json:"conversion_valuation_cap,omitempty"`
	// For cash proceeds calculation during a liquidity event.
	ExitMultiple *types.Ratio `json:"exit_multiple,omitempty"`
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
}
