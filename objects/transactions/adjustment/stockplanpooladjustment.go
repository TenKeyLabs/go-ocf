package adjustment

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing the change in the size of a Stock Plan pool.
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that affect a stock
// plan
type StockPlanPoolAdjustment struct {
	// Date on which board approved the change to the plan.
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// The number of shares reserved in the pool for this stock plan by the Board or equivalent
	// body as of the effective date of this pool adjustment.
	SharesReserved string `json:"shares_reserved"`
	// Identifier of the Stock Plan object, a subject of this transaction
	StockPlanID *string `json:"stock_plan_id"`
	// This optional field tracks when the stockholders approved this change to the stock plan.
	// This is intended for use by US companies that want to issue Incentive Stock Options
	// (ISOs), as the issuing StockPlan must receive shareholder approval within a specified
	// time frame in order to issue valid ISOs.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
