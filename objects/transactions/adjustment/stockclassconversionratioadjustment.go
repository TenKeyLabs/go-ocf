package adjustment

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/conversionmechanisms"
)

// Object describing the conversion ratio adjustment of a stock class that has a
// RatioConversionMechanism conversion mechanism where there was an actual repricing due to
// a down-round. The actual determination of the new conversion ratio / conversion price is
// calculated outside of OCF, so the specific mechanism - e.g. broad-based weighted-average
// anti-dilution protection vs. full ratchet anti-dilution protection.
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that affect the
// stock class
type StockClassConversionRatioAdjustment struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// New conversion ratio mechanism describing new conversion price and conversion ratio in
	// effect following a repricing - based on original issue price to new conversion price
	// (provided in this transaction). For 2-for-1 split the numerator of the ratio is 2 and the
	// denominator is 1.
	NewRatioConversionMechanism conversionmechanisms.RatioConversionMechanism `json:"new_ratio_conversion_mechanism"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Identifier of the StockClass object, a subject of this transaction
	StockClassID string `json:"stock_class_id"`
}
