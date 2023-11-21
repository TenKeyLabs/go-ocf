package objects

import "github.com/tenkeylabs/go-ocf/primitives/objects"

// Object describing a stock legend template
type StockLegendTemplate struct {
	objects.Object

	// Name for the stock legend template
	Name string `json:"name"`
	// The full text of the stock legend
	Text string `json:"text"`
}
