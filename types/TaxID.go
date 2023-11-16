package types

// Type representation of a government identifier for tax purposes (e.g. EIN) and
// corresponding country code (ISO-3166)
type TaxID struct {
	// Issuing country code (ISO 3166-1 alpha-2) for the tax identifier
	Country string `json:"country"`
	// Tax identifier as string
	TaxID string `json:"tax_id"`
}
