package enums

// Type of e-mail address (e.g. personal or business)
//
// Enumeration of email types
type EmailType string

const (
	EnumEmailTypeBUSINESS EmailType = "BUSINESS"
	EnumEmailTypeOTHER    EmailType = "OTHER"
	Personal              EmailType = "PERSONAL"
)
