package types

import "github.com/tenkeylabs/go-ocf/enums"

// Type representation of a termination window
type TerminationWindow struct {
	// The length of the period in this termination window (in number of periods of type
	// period_type)
	Period int64 `json:"period"`
	// The type of period being measured (e.g. days or month)
	PeriodType enums.PeriodType `json:"period_type"`
	// What cause of termination is this window for?
	Reason enums.TerminationWindowType `json:"reason"`
}
