package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a stock issuance transaction by the issuer and held by a stakeholder
type StockIssuance struct {
	issuance.Issuance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// The cost basis for this particular stock
	CostBasis *types.Monetary `json:"cost_basis,omitempty"`
	// Optional field to flag certain special types of issuances (like RSAs)
	IssuanceType *enums.StockIssuanceType `json:"issuance_type,omitempty"`
	// Number of shares issued to the stakeholder
	Quantity string `json:"quantity"`
	// Range(s) of the specific share numbers included in this issuance. This is different from
	// a certificate number you might include in the `custom_id` field or the `security_id`
	// created in this issuance. This field should be used where, for whatever reason, shares
	// are not fungible and you must track, with each issuance, *which* specific share numbers
	// are included in the issuance - e.g. share numbers 1 - 100 and 250-300.
	ShareNumbersIssued []types.ShareNumberRange `json:"share_numbers_issued,omitempty"`
	// The price per share paid for the stock by the holder
	SharePrice types.Monetary `json:"share_price"`
	// Identifier of the stock class for this stock issuance
	StockClassID string `json:"stock_class_id"`
	// List of stock legend ids that apply to this stock
	StockLegendIDS []string `json:"stock_legend_ids"`
	// Identifier of StockPlan the Stock was issued from (in the case of RSAs or Stock issued
	// from a plan).
	StockPlanID *string `json:"stock_plan_id,omitempty"`
	// Identifier of the VestingTerms to which this security is subject. If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
}
