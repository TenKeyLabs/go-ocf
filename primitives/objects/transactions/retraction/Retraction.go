package retraction

// Abstract object describing a security retraction transaction
type Retraction struct {
	// Reason for the retraction
	ReasonText string `json:"reason_text"`
}
