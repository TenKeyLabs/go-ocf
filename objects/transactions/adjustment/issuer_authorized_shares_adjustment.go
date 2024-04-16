package adjustment

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

// Object describing an event to change the number of authoried shares at the issuer level.
type IssuerAuthorizedSharesAdjustment struct {
	transactions.IssuerTransaction
	transactions.StockClassTransaction
	transactions.Transaction
	objects.Object

	// Date on which the board approved the change to the stock class
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// The new number of shares authorized for this stock class as of the event of this
	// transaction
	NewSharesAuthorized string `json:"new_shares_authorized"`
	// This optional field tracks when the stockholders approved the change to the stock class.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
