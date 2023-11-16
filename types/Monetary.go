package types

// Valued price per share
//
// Type representation of an amount of money in a specified currency
type Monetary struct {
	// Numeric amount of money
	Amount string `json:"amount"`
	// ISO 4217 currency code
	Currency string `json:"currency"`
}
