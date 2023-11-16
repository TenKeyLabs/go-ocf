package enums

// Enumeration types of valuation inputs that go into a formula - e.g. use a specified value
// (`FIXED`), a cap (`VALUATION_CAP`) or actual valuation (`ACTUAL`).
type ValuationBasedFormulaType string

const (
	Actual ValuationBasedFormulaType = "ACTUAL"
	Cap    ValuationBasedFormulaType = "CAP"
	Fixed  ValuationBasedFormulaType = "FIXED"
)
