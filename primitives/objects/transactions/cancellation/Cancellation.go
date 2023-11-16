package cancellation

// Abstract object describing fields common to all cancellation transaction objects
type Cancellation struct {
	// Identifier for the security that holds the remainder balance (for partial cancellations)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Reason for the cancellation
	ReasonText string `json:"reason_text"`
}
