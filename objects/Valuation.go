package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a valuation used in the cap table
type Valuation struct {
	objects.Object

	// Date on which board approved the valuation. This is essential for 409A valuations, in
	// particular, which require the Board to approve the valuation.
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Date on which this valuation is first valid
	EffectiveDate string `json:"effective_date"`
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
