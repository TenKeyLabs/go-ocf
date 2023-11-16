package types

import "github.com/tenkeylabs/go-ocf/enums"

// Type representation of a phone number
type Phone struct {
	// A valid phone number string in ITU E.123 international notation (e.g. +123 123 456 7890)
	PhoneNumber string `json:"phone_number"`
	// Type of phone number (e.g. mobile, home or business)
	PhoneType enums.PhoneType `json:"phone_type"`
}
