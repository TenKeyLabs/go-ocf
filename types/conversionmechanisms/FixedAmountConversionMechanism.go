package conversionmechanisms

import "github.com/tenkeylabs/go-ocf/enums"

// Describes how a security converts into a fixed amount of a stock class
//
// Abstract type setting forth required field(s) for ALL conversion mechanism types
type FixedAmountConversionMechanism struct {
	// How many shares of target Stock Class does this security convert into?
	ConvertsToQuantity string `json:"converts_to_quantity"`
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
}
