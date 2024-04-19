package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing securities issuance transaction by the issuer and held by a stakeholder
// as a form of compensation (as noted elsewhere, RSAs are not included here intentionally
// and should be modelled using Stock Issuances).
type EquityCompensationIssuance struct {
	issuance.Issuance
	transactions.SecurityTransaction

	BasePrice *types.Monetary `json:"base_price,omitempty"`
	// If the plan security is compensation, what kind?
	CompensationType enums.CompensationType `json:"compensation_type"`
	// Is this Equity Compensation exercisable prior to completion of vesting? If so, it's
	// assumed the vesting schedule will remain in effect but, instead of vesting a right to
	// exercise, it becomes the schedule determining when a right to repurchase the resulting
	// stock lapses.
	EarlyExercisable *bool `json:"early_exercisable,omitempty"`
	// If this is an option, what is the exercise price of the option?
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// Expiration date of the plan security
	ExpirationDate *string `json:"expiration_date"`
	// If the plan security is an option, what kind?
	OptionGrantType *enums.OptionType `json:"option_grant_type,omitempty"`
	// How many shares are subject to this plan security?
	Quantity string `json:"quantity"`
	// If the equity compensation was NOT issued from a plan (don't forget, plan-less options
	// are a thing), we need to know what underlying stock class it converts to.
	StockClassID *string `json:"stock_class_id,omitempty"`
	// If the equity compensation was issued from a plan (don't forget, plan-less options are a
	// thing), what is the plan id.
	StockPlanID *string `json:"stock_plan_id,omitempty"`
	// Exercise periods applicable to plan security after a termination for a given, enumerated
	// reason
	TerminationExerciseWindows []types.TerminationWindow `json:"termination_exercise_windows"`
	// Identifier of the VestingTerms to which this security is subject.  If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
}
