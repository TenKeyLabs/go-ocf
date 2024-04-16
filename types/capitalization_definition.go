package types

// If this conversion event and its `quantity_converted` value was based on the company's
// capitalization, please specify what stock classes, stock plans and securities were
// aggregated to calculate the capitalization value used in the calculation (e.g. if it was
// based on "fully diluted" capitalization, please provide details on how this was
// calculated using the capitalization type datastructure).
//
// Type represents a group of securities that constitutes some formally defined part of the
// company (e.g. post-money capitalization vs pre-money for a security)
type CapitalizationDefinition struct {
	// Securities (whether Stock, Plan Securities, Convertibles or Warrants) with these security
	// ids should be excluded from this definition of capitalization (overrides plan or
	// class-level rules)
	ExcludeSecurityIDS []string `json:"exclude_security_ids"`
	// Securities (whether Stock, Plan Securities, Convertibles or Warrants) with these security
	// ids should be included from this definition of capitalization (overrides plan or
	// class-level rules)
	IncludeSecurityIDS []string `json:"include_security_ids"`
	// All issuances of stock classes with these ids should be included (unless such an issuance
	// is specifically included in `exclude_security_ids`
	IncludeStockClassIDS []string `json:"include_stock_class_ids"`
	// All issuances of plan securities from stock plans with these ids should be included
	// (unless such an issuance is specifically excluded in `exclude_security_ids`
	IncludeStockPlansIDS []string `json:"include_stock_plans_ids"`
}
