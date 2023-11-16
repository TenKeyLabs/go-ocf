package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a stock issuance transaction by the issuer and held by a stakeholder
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all issuance objects
type StockIssuance struct {
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security issuance
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// The cost basis for this particular stock
	CostBasis *types.Monetary `json:"cost_basis,omitempty"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Optional field to flag certain special types of issuances (like RSAs)
	IssuanceType *enums.StockIssuanceType `json:"issuance_type,omitempty"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Number of shares issued to the stakeholder
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
	// Range(s) of the specific share numbers included in this issuance. This is different from
	// a certificate number you might include in the `custom_id` field or the `security_id`
	// created in this issuance. This field should be used where, for whatever reason, shares
	// are not fungible and you must track, with each issuance, *which* specific share numbers
	// are included in the issuance - e.g. share numbers 1 - 100 and 250-300.
	ShareNumbersIssued []types.ShareNumberRange `json:"share_numbers_issued,omitempty"`
	// The price per share paid for the stock by the holder
	SharePrice types.Monetary `json:"share_price"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID string `json:"stakeholder_id"`
	// Identifier of the stock class for this stock issuance
	StockClassID string `json:"stock_class_id"`
	// List of stock legend ids that apply to this stock
	StockLegendIDS []string `json:"stock_legend_ids"`
	// Identifier of StockPlan the Stock was issued from (in the case of RSAs or Stock issued
	// from a plan).
	StockPlanID *string `json:"stock_plan_id,omitempty"`
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// Identifier of the VestingTerms to which this security is subject. If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
}
