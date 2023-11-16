package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a stakeholder
//
// Abstract object to be extended by all other objects
type Stakeholder struct {
	// Addresses for the stakeholder
	Addresses []types.Address `json:"addresses,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// The contact info for an individual stakeholder
	ContactInfo *types.ContactInfoWithoutName `json:"contact_info,omitempty"`
	// What is the current relationship of the stakeholder to the issuer?
	CurrentRelationship *enums.StakeholderRelationshipType `json:"current_relationship,omitempty"`
	// Identifier for the object
	ID string `json:"id"`
	// This might be any sort of id assigned to the stakeholder by the issuer, such as an
	// internal company ID for an employee stakeholder
	IssuerAssignedID *string `json:"issuer_assigned_id,omitempty"`
	// Name for the stakeholder
	Name types.Name `json:"name"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// The primary contact info for an institutional stakeholder
	PrimaryContact *types.ContactInfo `json:"primary_contact,omitempty"`
	// Distinguish individuals from institutions
	StakeholderType enums.StakeholderType `json:"stakeholder_type"`
	// The tax ids for this stakeholder
	TaxIDS []types.TaxID `json:"tax_ids,omitempty"`
}
