package exercise

import "github.com/tenkeylabs/go-ocf/enums"

type Exercise struct {
	*EquityCompensationExercise
	*PlanSecurityExercise
	*WarrantExercise

	// This is done to avoid a breaking change as we work towards a bigger restructure of the
	// equity types in v2.0.0. `TX_PLAN_SECURITY_EXERCISE` will be deprecated in v2.0.0
	//
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security exercise
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Quantity of shares exercised
	Quantity string `json:"quantity"`
	// Identifier for the security (or securities) that resulted from the exercise
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
}
