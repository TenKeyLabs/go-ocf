package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a valuation used in the cap table
//
// Abstract object to be extended by all other objects
type Valuation struct {
	// Date on which board approved the valuation. This is essential for 409A valuations, in
	// particular, which require the Board to approve the valuation.
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which this valuation is first valid
	EffectiveDate string `json:"effective_date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Valued price per share
	PricePerShare types.Monetary `json:"price_per_share"`
	// Entity which provided the valuation
	Provider *string `json:"provider,omitempty"`
	// Identifier of the stock class for this valuation
	StockClassID string `json:"stock_class_id"`
	// This optional field tracks when the stockholders approved the valuation.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// Seam for supporting different types of valuations in future versions
	ValuationType enums.ValuationType `json:"valuation_type"`
}
