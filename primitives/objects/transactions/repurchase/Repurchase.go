package repurchase

import "github.com/tenkeylabs/go-ocf/types"

// Abstract object describing common properties to a repurchase transaction
type Repurchase struct {
	// Identifier for the security that holds the remainder balance (for partial repurchases)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Unstructured text description of consideration provided in exchange for security
	// repurchase
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Repurchase price per share of the stock
	Price types.Monetary `json:"price"`
	// Number of shares of stock repurchased
	Quantity string `json:"quantity"`
}
