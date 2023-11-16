package enums

// What type of interest compounding?
//
// Enumeration of interest compounding types
type CompoundingType string

const (
	Compounding CompoundingType = "COMPOUNDING"
	Simple      CompoundingType = "SIMPLE"
)
