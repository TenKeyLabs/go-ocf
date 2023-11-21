package adjustment

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/types/conversionmechanisms"
)

// Object describing the conversion ratio adjustment of a stock class that has a
// RatioConversionMechanism conversion mechanism where there was an actual repricing due to
// a down-round. The actual determination of the new conversion ratio / conversion price is
// calculated outside of OCF, so the specific mechanism - e.g. broad-based weighted-average
// anti-dilution protection vs. full ratchet anti-dilution protection.
type StockClassConversionRatioAdjustment struct {
	transactions.StockClassTransaction
	transactions.Transaction
	objects.Object

	// New conversion ratio mechanism describing new conversion price and conversion ratio in
	// effect following a repricing - based on original issue price to new conversion price
	// (provided in this transaction). For 2-for-1 split the numerator of the ratio is 2 and the
	// denominator is 1.
	NewRatioConversionMechanism conversionmechanisms.RatioConversionMechanism `json:"new_ratio_conversion_mechanism"`
}
