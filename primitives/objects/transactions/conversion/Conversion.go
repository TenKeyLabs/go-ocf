package conversion

// Abstract object describing fields common to all conversion transaction objects
type Conversion struct {
	// Identifier for the security (or securities) that resulted from the conversion
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
}
