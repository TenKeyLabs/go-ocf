package issuance

import "github.com/tenkeylabs/go-ocf/types"

// Abstract object describing fields common to all issuance objects
type Issuance struct {
	// Date of board approval for the security
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text description of consideration provided in exchange for security issuance
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// A custom ID for this security (e.g. CN-1.)
	CustomID string `json:"custom_id"`
	// List of security law exemptions (and applicable jurisdictions) for this security
	SecurityLawExemptions []types.SecurityExemption `json:"security_law_exemptions"`
	// Identifier for the stakeholder that holds legal title to this security
	StakeholderID string `json:"stakeholder_id"`
	// Date on which the stockholders approved the security
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
