package types

// Type representation of a primary contact person for a stakeholder (e.g. a fund)
type ContactInfo struct {
	// Emails to reach the contact at
	Emails []Email `json:"emails,omitempty"`
	// Contact's name
	Name Name `json:"name"`
	// Phone numbers to reach the contact at
	PhoneNumbers []Phone `json:"phone_numbers,omitempty"`
}
