package enums

// How should fractional shares be rounded?
//
// Enumeration of rounding types
type RoundingType string

const (
	Ceiling RoundingType = "CEILING"
	Floor   RoundingType = "FLOOR"
	Normal  RoundingType = "NORMAL"
)
