package enums

// Should the conversion amount be based on pre or post money capitalization
//
// Enumeration of convertible conversion timing for calculation purposes (e.g. does the
// instrument convert based on pre or post money).
type ConversionTimingType string

const (
	PostMoney ConversionTimingType = "POST_MONEY"
	PreMoney  ConversionTimingType = "PRE_MONEY"
)
