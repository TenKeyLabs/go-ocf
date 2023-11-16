package types

import "github.com/tenkeylabs/go-ocf/enums"

// Type representation of an email address
type Email struct {
	// A valid e-mail address
	EmailAddress string `json:"email_address"`
	// Type of e-mail address (e.g. personal or business)
	EmailType enums.EmailType `json:"email_type"`
}
