package adjustment

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing an event to change the number of authoried shares of a stock class.
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that affect the
// stock class
type StockClassAuthorizedSharesAdjustment struct {
	// Date on which the board approved the change to the stock class
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// The new number of shares authorized for this stock class as of the event of this
	// transaction
	NewSharesAuthorized string `json:"new_shares_authorized"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Identifier of the StockClass object, a subject of this transaction
	StockClassID string `json:"stock_class_id"`
	// This optional field tracks when the stockholders approved the change to the stock class.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
