package release

import "github.com/tenkeylabs/go-ocf/types"

// Abstract object describing fields common to all release transaction objects
type Release struct {
	// Unstructured text description of consideration provided in exchange for security release
	ConsiderationText *string `json:"consideration_text,omitempty"`
	// Quantity of shares released
	Quantity string `json:"quantity"`
	// The release price used to determine the value of the security at the time of release
	ReleasePrice types.Monetary `json:"release_price"`
	// Identifier of the new security (or securities) issuance resulting from a release
	// transaction
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// The settlement date for the shares released, typically after the release transaction date
	SettlementDate string `json:"settlement_date"`
}
