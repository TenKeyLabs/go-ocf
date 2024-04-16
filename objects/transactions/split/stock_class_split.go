package split

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a split of a stock class
type StockClassSplit struct {
	transactions.StockClassTransaction
	transactions.Transaction
	objects.Object

	// Ratio of new shares to old shares. For 2-for-1 split the numerator of the ratio is 2 and
	// the denominator is 1.
	SplitRatio *types.Ratio `json:"split_ratio,omitempty"`
}
