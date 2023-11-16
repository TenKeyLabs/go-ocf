package types

import "github.com/tenkeylabs/go-ocf/enums"

// Type representation of an address
type Address struct {
	// What type of address is this (e.g. legal address, contact address, etc.)
	AddressType enums.AddressType `json:"address_type"`
	// City
	City *string `json:"city,omitempty"`
	// Country code for this address (ISO 3166-1 alpha-2)
	Country string `json:"country"`
	// State, province, or equivalent identifier required for an address in this country
	CountrySubdivision *string `json:"country_subdivision,omitempty"`
	// Address postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// Street address (multi-line string)
	StreetSuite *string `json:"street_suite,omitempty"`
}
