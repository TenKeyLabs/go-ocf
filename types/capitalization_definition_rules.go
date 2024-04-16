package types

// The rules for which types of securities would be included in the capitalization
// definition.
//
// Type represents the rules for determining the capitalization definition for a security
type CapitalizationDefinitionRules struct {
	// Include the shares reserved for increasing option plans beyond the amount needed for any
	// promised options in the capitalization definition
	IncludeAdditionalOptionPoolTopup bool `json:"include_additional_option_pool_topup"`
	// Include the shares issued for any new share subscriptions that are part of the conversion
	// event in the capitalization definition
	IncludeNewMoney bool `json:"include_new_money"`
	// Include the shares reserved for increasing option plans to cover all promised options in
	// the capitalization definition
	IncludeOptionPoolTopupForPromisedOptions bool `json:"include_option_pool_topup_for_promised_options"`
	// Include the shares issued for converting all other convertibles that are converted as
	// part of the conversion event in the capitalization definition
	IncludeOtherConvertingSecurities bool `json:"include_other_converting_securities"`
	// Include all outstanding options in the capitalization definition
	IncludeOutstandingOptions bool `json:"include_outstanding_options"`
	// Include all outstanding share issuances in the capitalization definition
	IncludeOutstandingShares bool `json:"include_outstanding_shares"`
	// Include all outstanding options that have been reserved but have not been issued yet in
	// the capitalization definition
	IncludeOutstandingUnissuedOptions bool `json:"include_outstanding_unissued_options"`
	// Include the shares issued for converting this security in the capitalization definition
	IncludeThisSecurity bool `json:"include_this_security"`
}
