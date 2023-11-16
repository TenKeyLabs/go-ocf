package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing securities issuance transaction by the issuer and held by a stakeholder
// as a form of compensation (as noted elsewhere, RSAs are not included here intentionally
// and should be modelled using Stock Issuances).
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all issuance objects
type EquityCompensationIssuance struct {
	// If this is a stock appreciation right, what is the base price used to calculate the
	// appreciation of the SAR?
	BasePrice *types.Monetary `json:"base_price,omitempty"`
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// If the plan security is compensation, what kind?
	CompensationType enums.CompensationType `json:"compensation_type"`
	// Unstructured text description of consideration provided in exchange for security issuance
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Is this Equity Compensation exercisable prior to completion of vesting? If so, it's
	// assumed the vesting schedule will remain in effect but, instead of vesting a right to
	// exercise, it becomes the schedule determining when a right to repurchase the resulting
	// stock lapses.
	EarlyExercisable *bool `json:"early_exercisable,omitempty"`
	// If this is an option, what is the exercise price of the option?
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// Expiration date of the plan security
	ExpirationDate *string `json:"expiration_date"`
	// Identifier for the object
	ID string `json:"id"`
	// This is done to avoid a breaking change as we work towards a bigger restructure of the
	// equity types in v2.0.0. `TX_PLAN_SECURITY_ISSUANCE` will be deprecated in v2.0.0
	//
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// If the plan security is an option, what kind?
	OptionGrantType *enums.OptionType `json:"option_grant_type,omitempty"`
	// How many shares are subject to this plan security?
	Quantity string `json:"quantity"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// List of security law exemptions (and applicable jurisdictions) for this security
	SecurityLawExemptions []types.SecurityExemption `json:"security_law_exemptions"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID string `json:"stakeholder_id"`
	// If the equity compensation was NOT issued from a plan (don't forget, plan-less options
	// are a thing), we need to know what underlying stock class it converts to.
	StockClassID *string `json:"stock_class_id,omitempty"`
	// If the equity compensation was issued from a plan (don't forget, plan-less options are a
	// thing), what is the plan id.
	StockPlanID *string `json:"stock_plan_id,omitempty"`
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// Exercise periods applicable to plan security after a termination for a given, enumerated
	// reason
	TerminationExerciseWindows []types.TerminationWindow `json:"termination_exercise_windows"`
	// Identifier of the VestingTerms to which this security is subject.  If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
}
