package reissuance

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing a re-issuance of stock
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing common properties to a reissuance of a security
type StockReissuance struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Free-form human-readable reason for stock reissuance
	ReasonText *string `json:"reason_text,omitempty"`
	// Identifier of the new security (or securities) issuance resulting from a reissuance
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// When stock is reissued as a result of a stock split, this field contains id of the
	// respective stock class split transaction. It is not set otherwise.
	SplitTransactionID *string `json:"split_transaction_id,omitempty"`
}
