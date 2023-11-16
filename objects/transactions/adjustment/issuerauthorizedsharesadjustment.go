package adjustment

import (
	"github.com/tenkeylabs/go-ocf/enums"
)

type IssuerAuthorizedSharesAdjustment struct {
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
	// This optional field tracks when the stockholders approved the change to the stock class.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
