package transactions

// Abstract transaction object to be extended by all transaction objects that affect a stock
// plan
type StockPlanTransaction struct {
	// Identifier of the Stock Plan object, a subject of this transaction
	StockPlanID *string `json:"stock_plan_id"`
}
