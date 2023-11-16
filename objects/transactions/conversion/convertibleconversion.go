package conversion

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a conversion of a convertible security
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all conversion transaction objects
type ConvertibleConversion struct {
	// Identifier for the convertible that holds the remainder balance (for partial conversions)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// If this conversion event and its `quantity_converted` value was based on the company's
	// capitalization, please specify what stock classes, stock plans and securities were
	// aggregated to calculate the capitalization value used in the calculation (e.g. if it was
	// based on "fully diluted" capitalization, please provide details on how this was
	// calculated using the capitalization type datastructure).
	CapitalizationDefinition *types.CapitalizationDefinition `json:"capitalization_definition,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Quantity of security units converted
	QuantityConverted *string `json:"quantity_converted,omitempty"`
	// Reason for the conversion
	ReasonText string `json:"reason_text"`
	// Identifier for the security (or securities) that resulted from the conversion
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// What is the id of the convertible's conversion trigger that resulted in this conversion
	TriggerID string `json:"trigger_id"`
}
