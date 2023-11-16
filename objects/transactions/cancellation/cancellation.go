package cancellation

import "github.com/tenkeylabs/go-ocf/enums"

type Cancellation struct {
	*ConvertibleCancellation
	*EquityCompensationCancellation
	*PlanSecurityCancellation
	*StockCancellation
	*WarrantCancellation
	// Identifier for the security that holds the remainder balance (for partial cancellations)
	BalanceSecurityID *string `json:"balance_security_id,omitempty"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Quantity of non-monetary security units cancelled
	Quantity string `json:"quantity"`
	// Reason for the cancellation
	ReasonText string `json:"reason_text"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
}
