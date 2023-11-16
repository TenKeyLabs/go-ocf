package types

// One share of this stock class converts into this many target stock class shares
//
// Type representation of a ratio as two parts of a quotient, i.e. numerator and denominator
// numeric values
//
// For cash proceeds calculation during a liquidity event.
type Ratio struct {
	// Denominator of the ratio, i.e. the ratio of A to B (A:B) can be expressed as a fraction
	// (A/B), where B is the denominator
	Denominator string `json:"denominator,omitempty"`
	// Numerator of the ratio, i.e. the ratio of A to B (A:B) can be expressed as a fraction
	// (A/B), where A is the numerator
	Numerator string `json:"numerator,omitempty"`
}
