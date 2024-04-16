package enums

// If a security issued under this Stock Plan is cancelled, what happens to the reserved
// shares by default? NOTE: for any given security issued from the pool, the Plan's default
// cancellation behavior can be overridden by subsequent transactions cancelling the
// reserved stock, returning it to pool or marking it as capital stock. The event chain
// should always control - do not rely on this field and fail to traverse the events.
//
// For a given stock plan, what is the default rule for what happens to the shares reserved
// for a Plan Security after it's cancelled.
type StockPlanCancellationBehaviorType string

const (
	DefinedPerPlanSecurity StockPlanCancellationBehaviorType = "DEFINED_PER_PLAN_SECURITY"
	HoldAsCapitalStock     StockPlanCancellationBehaviorType = "HOLD_AS_CAPITAL_STOCK"
	Retire                 StockPlanCancellationBehaviorType = "RETIRE"
	ReturnToPool           StockPlanCancellationBehaviorType = "RETURN_TO_POOL"
)
