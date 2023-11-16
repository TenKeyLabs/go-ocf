package release

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing plan security release transaction (a compatibility wrapper for equity
// compensation release event
//
// # Object describing equity compensation security release transaction
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all release transaction objects
type PlanSecurityRelease struct {
	// This is done to avoid a breaking change as we work towards a bigger restructure of the
	// equity types in v2.0.0. `TX_PLAN_SECURITY_RELEASE` will be deprecated in v2.0.0
	//
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security release
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Quantity of shares released
	Quantity string `json:"quantity"`
	// The release price used to determine the value of the security at the time of release
	ReleasePrice types.Monetary `json:"release_price"`
	// Identifier of the new security (or securities) issuance resulting from a release
	// transaction
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// The settlement date for the shares released, typically after the release transaction date
	SettlementDate string `json:"settlement_date"`
}
