package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects/transactions/acceptance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/adjustment"
	"github.com/tenkeylabs/go-ocf/objects/transactions/cancellation"
	"github.com/tenkeylabs/go-ocf/objects/transactions/conversion"
	"github.com/tenkeylabs/go-ocf/objects/transactions/exercise"
	"github.com/tenkeylabs/go-ocf/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/reissuance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/release"
	"github.com/tenkeylabs/go-ocf/objects/transactions/repurchase"
	"github.com/tenkeylabs/go-ocf/objects/transactions/retraction"
	"github.com/tenkeylabs/go-ocf/objects/transactions/returntopool"
	"github.com/tenkeylabs/go-ocf/objects/transactions/split"
	"github.com/tenkeylabs/go-ocf/objects/transactions/transfer"
	"github.com/tenkeylabs/go-ocf/objects/transactions/vesting"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversionmechanisms"
	"github.com/tenkeylabs/go-ocf/types/conversionrights"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

type Transaction struct {
	*acceptance.Acceptance
	*cancellation.Cancellation
	*conversion.Conversion
	*exercise.Exercise
	*issuance.Issuance
	*reissuance.StockReissuance
	*repurchase.StockRepurchase
	*release.EquityCompensationRelease
	*retraction.Retraction
	*returntopool.StockPlanReturnToPool
	*split.StockClassSplit
	*adjustment.Adjustment
	*transfer.Transfer
	*vesting.Vesting
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID *string `json:"security_id,omitempty"`
	// Identifier for the security that holds the remainder balance (for partial cancellations)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Reason for the cancellation
	ReasonText *string `json:"reason_text,omitempty"`
	// Identifier for the security (or securities) that resulted from the conversion
	ResultingSecurityIDS []string `json:"resulting_security_ids,omitempty"`
	// What is the id of the convertible's conversion trigger that resulted in this conversion
	TriggerID *string `json:"trigger_id,omitempty"`
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// In event the convertible can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ConversionTriggers []conversiontriggers.ConversionTrigger `json:"conversion_triggers,omitempty"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID *string `json:"custom_id,omitempty"`
	// List of security law exemptions (and applicable jurisdictions) for this security
	SecurityLawExemptions []types.SecurityExemption `json:"security_law_exemptions,omitempty"`
	// If different convertible instruments have seniorty over one another, use this value to
	// build a seniority stack, with 1 being highest seniority and equal seniority values
	// assumed to be equal priority
	Seniority *int64 `json:"seniority,omitempty"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID *string `json:"stakeholder_id,omitempty"`
	// Unstructured text description of consideration provided in exchange for security exercise
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Quantity of shares exercised
	Quantity *string `json:"quantity,omitempty"`
	// If the plan security is compensation, what kind?
	CompensationType *enums.CompensationType `json:"compensation_type,omitempty"`
	// If this is an option, what is the exercise price of the option?
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// Expiration date of the plan security
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// Exercise periods applicable to plan security after a termination for a given, enumerated
	// reason
	TerminationExerciseWindows []types.TerminationWindow `json:"termination_exercise_windows,omitempty"`
	// Identifier of the VestingTerms to which this security is subject.  If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
	// What conversion mechanism applies to calculate the number of resulting securities?
	ConversionMechanism *conversionmechanisms.ConversionMechanism `json:"conversion_mechanism,omitempty"`
	// When the conditions of the trigger are met, how does the convertible convert?
	ConversionRight *conversionrights.ConversionRights `json:"conversion_right,omitempty"`
	// Human-friendly nickname to describe the conversion right
	Nickname *string `json:"nickname,omitempty"`
	// Legal language describing what conditions must be satisfied for the conversion to take
	// place (ideally, this should be excerpted from the instrument where possible)
	TriggerCondition *string `json:"trigger_condition,omitempty"`
	// Long-form description of the trigger
	TriggerDescription *string `json:"trigger_description,omitempty"`
	// When the trigger condition is met, is the conversion automatic, elective or automatic
	// with an elective right not to convert
	Type *enums.ConversionTriggerType `json:"type,omitempty"`
	// Is this stock class potentially convertible into a future, as-yet undetermined stock
	// class (e.g. Founder Preferred)
	ConvertsToFutureRound *bool `json:"converts_to_future_round,omitempty"`
	// Is this an MFN (Most Favored Nations) flavored Convertible Note?
	ConversionMfn *bool `json:"conversion_mfn,omitempty"`
	// Reference to the `id` of a VestingCondition in this security's VestingTerms. This
	// condition should have a trigger type of `VESTING_EVENT`.
	VestingConditionID *string `json:"vesting_condition_id,omitempty"`
	// If the plan security is an option, what kind?
	OptionGrantType *enums.OptionType `json:"option_grant_type,omitempty"`
	// The identifier of the existing, known stock class this stock class can convert into
	ConvertsToStockClassID *string `json:"converts_to_stock_class_id,omitempty"`
	// If the equity compensation was issued from a plan (don't forget, plan-less options are a
	// thing), what is the plan id.
	StockPlanID *string `json:"stock_plan_id,omitempty"`
}
