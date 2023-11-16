package types

// Type representation of the contact info for an individual stakeholder
type ContactInfoWithoutName struct {
	// Emails to reach the contact at
	Emails []Email `json:"emails,omitempty"`
	// Phone numbers to reach the contact at
	PhoneNumbers []Phone `json:"phone_numbers,omitempty"`
}
