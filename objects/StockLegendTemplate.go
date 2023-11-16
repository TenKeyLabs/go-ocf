package objects

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing a stock legend template
//
// Abstract object to be extended by all other objects
type StockLegendTemplate struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Identifier for the object
	ID string `json:"id"`
	// Name for the stock legend template
	Name string `json:"name"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// The full text of the stock legend
	Text string `json:"text"`
}
