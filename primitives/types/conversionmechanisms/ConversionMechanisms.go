package conversionmechanisms

import "github.com/tenkeylabs/go-ocf/enums"

// Abstract type setting forth required field(s) for ALL conversion mechanism types
type ConversionMechanism struct {
	// Identifies the specific conversion trigger type
	Type enums.ConversionMechanismType `json:"type"`
}
