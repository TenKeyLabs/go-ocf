package transfer

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing plan security transfer transaction (a compatibility wrapper for equity
// compensation transfer event)
//
// # Object describing a transfer of equity compensation
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing a security transfer or secondary sale transaction
type PlanSecurityTransfer struct {
	// This is done to avoid a breaking change as we work towards a bigger restructure of the
	// equity types in v2.0.0. `TX_PLAN_SECURITY_TRANSFER` will be deprecated in v2.0.0
	//
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Identifier for the security that holds the remainder balance (for partial transfers)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security transfer
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Quantity of non-monetary security units transferred
	Quantity string `json:"quantity"`
	// Array of identifiers for new security (or securities) created as a result of the
	// transaction
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
}
