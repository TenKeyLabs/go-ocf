package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing convertible instrument issuance transaction by the issuer and held by a
// stakeholder
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract object describing fields common to all issuance objects
type ConvertibleIssuance struct {
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Unstructured text description of consideration provided in exchange for security issuance
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// In event the convertible can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ConversionTriggers []conversiontriggers.ConversionTrigger `json:"conversion_triggers"`
	// What kind of convertible instrument is this (of the supported, enumerated types)
	ConvertibleType enums.ConvertibleType `json:"convertible_type"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Amount invested and outstanding on date of issuance of this convertible
	InvestmentAmount types.Monetary `json:"investment_amount"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// What pro-rata (if any) is the holder entitled to buy at the next round?
	ProRata *string `json:"pro_rata,omitempty"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// List of security law exemptions (and applicable jurisdictions) for this security
	SecurityLawExemptions []types.SecurityExemption `json:"security_law_exemptions"`
	// If different convertible instruments have seniorty over one another, use this value to
	// build a seniority stack, with 1 being highest seniority and equal seniority values
	// assumed to be equal priority
	Seniority int64 `json:"seniority"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID string `json:"stakeholder_id"`
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
