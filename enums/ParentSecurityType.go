package enums

// Enumeration of parent sources a stock can be issued or created from
type ParentSecurityType string

const (
	Convertible ParentSecurityType = "CONVERTIBLE"
	Stock       ParentSecurityType = "STOCK"
	StockPlan   ParentSecurityType = "STOCK_PLAN"
	Warrant     ParentSecurityType = "WARRANT"
)
