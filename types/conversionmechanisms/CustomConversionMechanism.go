package conversionmechanisms

import "github.com/tenkeylabs/go-ocf/enums"

// Sets forth inputs and conversion mechanism of a custom conversion, a conversion type that
// cannot be accurately modelled with any other OCF conversion mechanism type
//
// Abstract type setting forth required field(s) for ALL conversion mechanism types
type CustomConversionMechanism struct {
	// Detailed description of how the number of resulting shares should be determined? Use
	// legal language from an instrument where possible
	CustomConversionDescription string `json:"custom_conversion_description"`
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
}
