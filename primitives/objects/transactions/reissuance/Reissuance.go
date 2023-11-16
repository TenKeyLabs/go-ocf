package reissuance

// Abstract object describing common properties to a reissuance of a security
type Reissuance struct {
	// Free-form human-readable reason for stock reissuance
	ReasonText *string `json:"reason_text,omitempty"`
	// Identifier of the new security (or securities) issuance resulting from a reissuance
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// When stock is reissued as a result of a stock split, this field contains id of the
	// respective stock class split transaction. It is not set otherwise.
	SplitTransactionID *string `json:"split_transaction_id,omitempty"`
}
