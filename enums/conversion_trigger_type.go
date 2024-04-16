package enums

// When the trigger condition is met, is the conversion automatic, elective or automatic
// with an elective right not to convert
//
// Enumeration of types of triggers common to various legal rights - e.g. does the
// satisfaction of a condition trigger an automatic conversion or merely a right to convert?
// If `UNSPECIFIED`, the system of record cannot represent this data in a structured form.
type ConversionTriggerType string

const (
	AutomaticOnCondition ConversionTriggerType = "AUTOMATIC_ON_CONDITION"
	AutomaticOnDate      ConversionTriggerType = "AUTOMATIC_ON_DATE"
	ElectiveAtWill       ConversionTriggerType = "ELECTIVE_AT_WILL"
	ElectiveInRange      ConversionTriggerType = "ELECTIVE_IN_RANGE"
	ElectiveOnCondition  ConversionTriggerType = "ELECTIVE_ON_CONDITION"
	Unspecified          ConversionTriggerType = "UNSPECIFIED"
)
