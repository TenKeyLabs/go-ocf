package enums

// If the plan security is an option, what kind?
//
// Enumeration of option types
type OptionType string

const (
	ISO  OptionType = "ISO"
	Intl OptionType = "INTL"
	Nso  OptionType = "NSO"
)
