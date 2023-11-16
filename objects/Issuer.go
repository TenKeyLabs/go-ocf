package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing the issuer of the cap table (the company whose cap table this is)
//
// Abstract object to be extended by all other objects
type Issuer struct {
	// The headquarters address of the issuing company
	Address *types.Address `json:"address,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// The country where the issuer company was legally formed (ISO 3166-1 alpha-2)
	CountryOfFormation string `json:"country_of_formation"`
	// The state, province, or subdivision where the issuer company was legally formed
	CountrySubdivisionOfFormation *string `json:"country_subdivision_of_formation,omitempty"`
	// Doing Business As name
	DBA *string `json:"dba,omitempty"`
	// A work email that the issuer company can be reached at
	Email *types.Email `json:"email,omitempty"`
	// Date of formation
	FormationDate string `json:"formation_date"`
	// Identifier for the object
	ID string `json:"id"`
	// The initial number of shares authorized for this issuer
	InitialSharesAuthorized *string `json:"initial_shares_authorized,omitempty"`
	// Legal name of the issuer
	LegalName string `json:"legal_name"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// A phone number that the issuer company can be reached at
	Phone *types.Phone `json:"phone,omitempty"`
	// The tax ids for this issuer company
	TaxIDS []types.TaxID `json:"tax_ids,omitempty"`
}
