package split

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a split of a stock class
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that affect the
// stock class
type StockClassSplit struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Ratio of new shares to old shares. For 2-for-1 split the numerator of the ratio is 2 and
	// the denominator is 1.
	SplitRatio *types.Ratio `json:"split_ratio,omitempty"`
	// Identifier of the StockClass object, a subject of this transaction
	StockClassID string `json:"stock_class_id"`
}
