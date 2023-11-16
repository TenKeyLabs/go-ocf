package enums

// The type of this stock class (e.g. Preferred or Common)
//
// Enumeration of stock class types
type StockClassType string

const (
	Common    StockClassType = "COMMON"
	Preferred StockClassType = "PREFERRED"
)
