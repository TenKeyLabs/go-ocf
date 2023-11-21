package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a stakeholder
type Stakeholder struct {
	objects.Object

	// Addresses for the stakeholder
	Addresses []types.Address `json:"addresses,omitempty"`
	// The contact info for an individual stakeholder
	ContactInfo *types.ContactInfoWithoutName `json:"contact_info,omitempty"`
	// What is the current relationship of the stakeholder to the issuer?
	CurrentRelationship *enums.StakeholderRelationshipType `json:"current_relationship,omitempty"`
	// This might be any sort of id assigned to the stakeholder by the issuer, such as an
	// internal company ID for an employee stakeholder
	IssuerAssignedID *string `json:"issuer_assigned_id,omitempty"`
	// Name for the stakeholder
	Name types.Name `json:"name"`
	// The primary contact info for an institutional stakeholder
	PrimaryContact *types.ContactInfo `json:"primary_contact,omitempty"`
	// Distinguish individuals from institutions
	StakeholderType enums.StakeholderType `json:"stakeholder_type"`
	// The tax ids for this stakeholder
	TaxIDS []types.TaxID `json:"tax_ids,omitempty"`
}
