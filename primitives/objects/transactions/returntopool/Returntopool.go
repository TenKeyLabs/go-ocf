package returntopool

// Abstract object describing a terminal transaction where securities return to a stock plan
// pool
type ReturnToPool struct {
	// How many shares were returned to the pool?
	Quantity string `json:"quantity"`
	// Reason for the return to the pool
	ReasonText string `json:"reason_text"`
	// Id of the Stock Plan whose pool the reserved shares should return to. This does not have
	// to be the same pool the securities were issued from as sometimes plan rollovers or other
	// actions taken by the company can result in stock returning to a different pool.
	StockPlanID string `json:"stock_plan_id"`
}
