package enums

// What kind of convertible instrument is this (of the supported, enumerated types)
//
// Enumeration of convertible instrument types
type ConvertibleType string

const (
	ConvertibleSecurity ConvertibleType = "CONVERTIBLE_SECURITY"
	Note                ConvertibleType = "NOTE"
	Safe                ConvertibleType = "SAFE"
)
