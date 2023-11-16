package enums

// What kind of conversion right is this?
//
// Enumeration of types of conversion rights.
type ConversionRightType string

const (
	ConvertibleConversionRight ConversionRightType = "CONVERTIBLE_CONVERSION_RIGHT"
	StockClassConversionRight  ConversionRightType = "STOCK_CLASS_CONVERSION_RIGHT"
	WarrantConversionRight     ConversionRightType = "WARRANT_CONVERSION_RIGHT"
)
