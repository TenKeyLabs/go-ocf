package transactions

// Abstract transaction object to be extended by all transaction objects that affect the
// stock class
type StockClassTransaction struct {
	// Identifier of the StockClass object, a subject of this transaction
	StockClassID string `json:"stock_class_id"`
}
