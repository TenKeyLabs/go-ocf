package conversion

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/conversion"
)

// Object describing a conversion of stock
type StockConversion struct {
	conversion.Conversion
	transactions.SecurityTransaction

	// Identifier for the security that holds the remainder balance (for partial conversions)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Quantity of non-monetary security units converted
	QuantityConverted string `json:"quantity_converted"`
}
