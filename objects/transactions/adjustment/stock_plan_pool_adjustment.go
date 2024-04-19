package adjustment

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

// Object describing the change in the size of a Stock Plan pool.
type StockPlanPoolAdjustment struct {
	transactions.StockPlanTransaction

	// Date on which board approved the change to the plan.
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// The number of shares reserved in the pool for this stock plan by the Board or equivalent
	// body as of the effective date of this pool adjustment.
	SharesReserved string `json:"shares_reserved"`
	// This optional field tracks when the stockholders approved this change to the stock plan.
	// This is intended for use by US companies that want to issue Incentive Stock Options
	// (ISOs), as the issuing StockPlan must receive shareholder approval within a specified
	// time frame in order to issue valid ISOs.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
