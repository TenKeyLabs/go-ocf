package transactions

// Abstract transaction object to be extended by all transaction objects that affect the
// issuer
type IssuerTransaction struct {
	// Identifier of the Issuer object, a subject of this transaction
	IssuerID string `json:"issuer_id,omitempty"`
}
