package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism based on price per share of a future round
// (with potential discounts)
type SharePriceBasedConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// A description of the specifics of the conversion - e.g. The Holder is entitled, during
	// the Exercise Period, to purchase from the Company such number of Preferred Shares as are
	// equal to $100,000 divided by the Exercise Price. 'Exercise Price' shall mean 80% of the
	// price per share paid by the investors in the next Qualified Financing.
	Description string `json:"description"`
	// True if the conversion shares should be based on a discount off the price-per-share in
	// the next elligible financing
	Discount *bool `json:"discount,omitempty"`
	// If the resulting conversion shares is based on a fixed amount discount off the
	// price-per-share of the next eilligible financing, what is the discount amount (in
	// currency)
	DiscountAmount *types.Monetary `json:"discount_amount,omitempty"`
	// If the conversion price is base on a percent discount off the price-per-share of the next
	// elligible financing, what is the discount percent
	DiscountPercentage *string `json:"discount_percentage,omitempty"`
}
