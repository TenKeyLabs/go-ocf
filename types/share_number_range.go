package types

// Type representation of a range of share numbers associated with an event (such as the
// share numbers associated with an issuance) - for use where shares are not fungible and
// need unique identifiers *per share*
type ShareNumberRange struct {
	// The ending share number of a range of shares impacted by a particular event
	// (**INCLUSIVE**)
	EndingShareNumber string `json:"ending_share_number"`
	// The starting share number of a range of shares impacted by a particular event
	// (**INCLUSIVE** and assuming **share counts start at 1**)
	StartingShareNumber string `json:"starting_share_number"`
}
