package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

// Object describing an acceleration of vesting, in which additional shares vest ahead of
// the schedule specified in security's vesting terms.
type VestingAcceleration struct {
	transactions.SecurityTransaction

	// Number of shares vesting ahead of schedule
	Quantity string `json:"quantity"`
	// Reason for the acceleration; unstructured text
	ReasonText string `json:"reason_text"`
}
