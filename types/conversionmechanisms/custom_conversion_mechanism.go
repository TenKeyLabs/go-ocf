package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"
)

// Sets forth inputs and conversion mechanism of a custom conversion, a conversion type that
// cannot be accurately modelled with any other OCF conversion mechanism type
type CustomConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// Detailed description of how the number of resulting shares should be determined? Use
	// legal language from an instrument where possible
	CustomConversionDescription string `json:"custom_conversion_description"`
}
