package returntopool

import "github.com/tenkeylabs/go-ocf/enums"

// Object describing which stock plan pool a particular security's shares were returned to
// upon cancellation.
//
// # Abstract object to be extended by all other objects
//
// # Abstract transaction object to be extended by all other transaction objects
//
// Abstract transaction object to be extended by all transaction objects that deal with
// individual securities
//
// Abstract transaction object to be extended by all transaction objects that affect a stock
// plan
//
// Abstract object describing a terminal transaction where securities return to a stock plan
// pool
type StockPlanReturnToPool struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Date on which the transaction occurred
	Date string `json:"date"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// How many shares were returned to the pool?
	Quantity string `json:"quantity"`
	// Reason for the return to the pool
	ReasonText string `json:"reason_text"`
	// Identifier for the security (stock, plan security, warrant, or convertible) by which it
	// can be referenced by other transaction objects. Note that while this identifier is
	// created with an issuance object, it should be different than the issuance object's `id`
	// field which identifies the issuance transaction object itself. All future transactions on
	// the security (e.g. acceptance, transfer, cancel, etc.) must reference this `security_id`
	// to qualify which security the transaction applies to.
	SecurityID string `json:"security_id"`
	// Identifier of the Stock Plan object, a subject of this transaction
	//
	// Id of the Stock Plan whose pool the reserved shares should return to. This does not have
	// to be the same pool the securities were issued from as sometimes plan rollovers or other
	// actions taken by the company can result in stock returning to a different pool.
	StockPlanID *string `json:"stock_plan_id"`
}
