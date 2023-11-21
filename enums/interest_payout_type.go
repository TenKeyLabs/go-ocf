package enums

// How is interest paid out (if at applicable)
//
// Enumeration of interest payout types (e.g. deferred or cash payment)
type InterestPayoutType string

const (
	Cash     InterestPayoutType = "CASH"
	Deferred InterestPayoutType = "DEFERRED"
)
