package conversion

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/conversion"
	"github.com/tenkeylabs/go-ocf/types"
)

// Object describing a conversion of a convertible security
type ConvertibleConversion struct {
	conversion.Conversion
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Identifier for the convertible that holds the remainder balance (for partial conversions)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// If this conversion event and its `quantity_converted` value was based on the company's
	// capitalization, please specify what stock classes, stock plans and securities were
	// aggregated to calculate the capitalization value used in the calculation (e.g. if it was
	// based on "fully diluted" capitalization, please provide details on how this was
	// calculated using the capitalization type datastructure).
	CapitalizationDefinition *types.CapitalizationDefinition `json:"capitalization_definition,omitempty"`
	// Quantity of security units converted
	QuantityConverted *string `json:"quantity_converted,omitempty"`
	// Reason for the conversion
	ReasonText string `json:"reason_text"`
	// What is the id of the convertible's conversion trigger that resulted in this conversion
	TriggerID string `json:"trigger_id"`
}
