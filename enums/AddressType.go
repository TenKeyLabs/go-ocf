package enums

// What type of address is this (e.g. legal address, contact address, etc.)
//
// Enumeration of address types
type AddressType string

const (
	Contact              AddressType = "CONTACT"
	EnumAddressTypeOTHER AddressType = "OTHER"
	Legal                AddressType = "LEGAL"
)
