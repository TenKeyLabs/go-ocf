package conversion

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

type Conversion struct {
	*ConvertibleConversion
	*StockConversion

	// In event the convertible can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ConversionTriggers []conversiontriggers.ConversionTrigger `json:"conversion_triggers"`
	// Identifier for the security that holds the remainder balance (for partial conversions)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Quantity of non-monetary security units converted
	QuantityConverted string `json:"quantity_converted"`
	// Identifier for the security (or securities) that resulted from the conversion
	ResultingSecurityIDS []string `json:"resulting_security_ids"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// What is the id of the convertible's conversion trigger that resulted in this conversion
	TriggerID string `json:"trigger_id"`
}
