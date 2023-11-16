package transfer

// Abstract object describing a security transfer or secondary sale transaction
type Transfer struct {
	// Identifier for the security that holds the remainder balance (for partial transfers)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Unstructured text description of consideration provided in exchange for security transfer
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Array of identifiers for new security (or securities) created as a result of the
	// transaction
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
}
