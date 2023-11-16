package transactions

// Abstract transaction object to be extended by all other transaction objects
type Transaction struct {
	// Date on which the transaction occurred
	Date string `json:"date"`
}
