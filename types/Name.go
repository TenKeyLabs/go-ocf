package types

// Type comprising of multiple name components
//
// Contact's name
type Name struct {
	// First/given name for the individual
	FirstName *string `json:"first_name,omitempty"`
	// Last/family name for the individual
	LastName *string `json:"last_name,omitempty"`
	// Legal full name for the individual/institution
	LegalName string `json:"legal_name"`
}
