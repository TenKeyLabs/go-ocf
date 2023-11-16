package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

// Object describing warrant issuance transaction by the issuer and held by a stakeholder
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all issuance objects
type WarrantIssuance struct {
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security issuance
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// The exercise price of the warrant
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// In event the Warrant can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ExerciseTriggers []conversiontriggers.ConversionTrigger `json:"exercise_triggers"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Actual purchase price of the warrant (sum up purported value of all consideration,
	// including in-kind)
	PurchasePrice types.Monetary `json:"purchase_price"`
	// Quantity of shares the warrant is exercisable for
	Quantity *string `json:"quantity,omitempty"`
	// If quantity is provided, use this to specify where the number came from - e.g. was it a
	// fixed value from the instrument (`INSTRUMENT_FIXED`), a human estimate
	// (`HUMAN_ESTIMATED`), etc. If quantity is provided and this field is not, this is assumed
	// to be `UNSPECIFIED`
	QuantitySource *enums.QuantitySourceType `json:"quantity_source,omitempty"`
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
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// Identifier of the VestingTerms to which this security is subject. If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
	// What is expiration date of the warrant (if applicable)
	WarrantExpirationDate *string `json:"warrant_expiration_date,omitempty"`
}
