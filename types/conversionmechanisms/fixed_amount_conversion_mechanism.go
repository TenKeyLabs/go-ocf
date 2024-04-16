package conversionmechanisms

import "github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"

// Describes how a security converts into a fixed amount of a stock class
type FixedAmountConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// How many shares of target Stock Class does this security convert into?
	ConvertsToQuantity string `json:"converts_to_quantity"`
}
