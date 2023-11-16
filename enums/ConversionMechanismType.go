package enums

// Identifies the specific conversion trigger type
//
// Enumeration of convertible conversion calculation types.
type ConversionMechanismType string

const (
	FixedAmountConversion                  ConversionMechanismType = "FIXED_AMOUNT_CONVERSION"
	FixedPercentOfCapitalizationConversion ConversionMechanismType = "FIXED_PERCENT_OF_CAPITALIZATION_CONVERSION"
	RatioConversion                        ConversionMechanismType = "RATIO_CONVERSION"
	SAFEConversion                         ConversionMechanismType = "SAFE_CONVERSION"
	ValuationBasedConversion               ConversionMechanismType = "VALUATION_BASED_CONVERSION"
	ConvertibleNoteConversion              ConversionMechanismType = "CONVERTIBLE_NOTE_CONVERSION"
	CustomConversion                       ConversionMechanismType = "CUSTOM_CONVERSION"
	PpsBasedConversion                     ConversionMechanismType = "PPS_BASED_CONVERSION"
)
