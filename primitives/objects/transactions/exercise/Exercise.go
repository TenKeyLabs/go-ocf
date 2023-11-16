package exercise

// Abstract object describing fields common to all exercise transaction objects
type Exercise struct {
	// Unstructured text description of consideration provided in exchange for security exercise
	ConsiderationText                                                                           *string  `json:"consideration_text,omitempty"`
	// Identifier for the security (or securities) that resulted from the exercise
	ResultingSecurityIDS                                                                        []string `json:"resulting_security_ids"`
}
