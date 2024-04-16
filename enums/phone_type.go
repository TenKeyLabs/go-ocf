package enums

// Type of phone number (e.g. mobile, home or business)
//
// Enumeration of phone number types
type PhoneType string

const (
	EnumPhoneTypeBUSINESS PhoneType = "BUSINESS"
	EnumPhoneTypeOTHER    PhoneType = "OTHER"
	Home                  PhoneType = "HOME"
	Mobile                PhoneType = "MOBILE"
)
