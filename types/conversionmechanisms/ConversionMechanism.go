package conversionmechanisms

import "github.com/tenkeylabs/go-ocf/enums"

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
}
