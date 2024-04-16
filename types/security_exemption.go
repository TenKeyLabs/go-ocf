package types

// Type representation of a securities issuance exemption that includes an unstructured
// description and a country code for ease of processing and analysis
type SecurityExemption struct {
	// Description of an applicable security law exemption governing the issuance
	Description                                                                                string `json:"description"`
	// Jurisdiction of the applicable law. This is a free-text field as there is no known
	// enumeration of all global legal jurisdictions, but please try to use ISO 3166-1 alpha-2,
	// if appropriate. Otherwise, we rely on implementers to choose an appropriate value here.
	Jurisdiction                                                                               string `json:"jurisdiction"`
}
