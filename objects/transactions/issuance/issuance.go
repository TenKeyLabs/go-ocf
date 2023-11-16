package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

type Issuance struct {
	*ConvertibleIssuance
	*EquityCompensationIssuance
	*PlanSecurityIssuance
	*StockIssuance
	*WarrantIssuance

	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// In event the convertible can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ConversionTriggers []conversiontriggers.ConversionTrigger `json:"conversion_triggers"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// List of security law exemptions (and applicable jurisdictions) for this security
	SecurityLawExemptions []types.SecurityExemption `json:"security_law_exemptions"`
	// If different convertible instruments have seniorty over one another, use this value to
	// build a seniority stack, with 1 being highest seniority and equal seniority values
	// assumed to be equal priority
	Seniority int64 `json:"seniority"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID string `json:"stakeholder_id"`
	// If the plan security is compensation, what kind?
	CompensationType enums.CompensationType `json:"compensation_type"`
	// If this is an option, what is the exercise price of the option?
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// Expiration date of the plan security
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// Exercise periods applicable to plan security after a termination for a given, enumerated
	// reason
	TerminationExerciseWindows []types.TerminationWindow `json:"termination_exercise_windows,omitmepty"`
	// Identifier of the VestingTerms to which this security is subject.  If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
	// If the plan security is an option, what kind?
	OptionGrantType *enums.OptionType `json:"option_grant_type,omitempty"`
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// If the equity compensation was issued from a plan (don't forget, plan-less options are a
	// thing), what is the plan id.
	StockPlanID *string `json:"stock_plan_id,omitempty"`
}
