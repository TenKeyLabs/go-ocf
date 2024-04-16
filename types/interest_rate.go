package types

// Type representation of an interest rate, including accrual start and end dates
type InterestRate struct {
	// Optional end date (inclusive) for interest accruing at the specified rate. If none
	// specified, interest will accrue indefinitely or until accrual of next interest rate
	// commences
	AccrualEndDate *string `json:"accrual_end_date,omitempty"`
	// Commencement date for interest accruing at the specified rate
	AccrualStartDate string `json:"accrual_start_date"`
	// Interest rate for the convertible (decimal representation - e.g. 0.125 for 12.5%)
	Rate string `json:"rate"`
}
