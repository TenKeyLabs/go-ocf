package exercise

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing a warrant exercise transaction
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all exercise transaction objects
type WarrantExercise struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security exercise
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Identifier for the security (or securities) that resulted from the exercise
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// What is the id of the warrant's exercise trigger that resulted in this exercise
	TriggerID string `json:"trigger_id"`
}
