package objects

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing a plan which stock options are issued from
//
// Abstract object to be extended by all other objects
type StockPlan struct {
	// Date on which board approved the plan
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// If a security issued under this Stock Plan is cancelled, what happens to the reserved
	// shares by default? NOTE: for any given security issued from the pool, the Plan's default
	// cancellation behavior can be overridden by subsequent transactions cancelling the
	// reserved stock, returning it to pool or marking it as capital stock. The event chain
	// should always control - do not rely on this field and fail to traverse the events.
	DefaultCancellationBehavior *enums.StockPlanCancellationBehaviorType `json:"default_cancellation_behavior,omitempty"`
	// Identifier for the object
	ID string `json:"id"`
	// The initial number of shares reserved in the pool for this stock plan by the Board or
	// equivalent body.
	InitialSharesReserved string `json:"initial_shares_reserved"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Name for the stock plan
	PlanName string `json:"plan_name"`
	// Identifier of the StockClass object this plan is composed of
	StockClassID string `json:"stock_class_id"`
	// This optional field tracks when the stockholders approved this stock plan. This is
	// intended for use by US companies that want to issue Incentive Stock Options (ISOs), as
	// the issuing StockPlan must receive shareholder approval within a specified time frame in
	// order to issue valid ISOs.
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
}
