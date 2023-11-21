package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism of a convertible note
type NoteConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// How is company capitalization defined for purposes of conversion? If possible, include
	// the legal language from the instrument.
	CapitalizationDefinition *string `json:"capitalization_definition,omitempty"`
	// The rules for which types of securities would be included in the capitalization
	// definition.
	CapitalizationDefinitionRules *types.CapitalizationDefinitionRules `json:"capitalization_definition_rules,omitempty"`
	// What type of interest compounding?
	CompoundingType enums.CompoundingType `json:"compounding_type"`
	// What is the percentage discount available upon conversion, if applicable? (decimal
	// representation - e.g. 0.125 for 12.5%)
	ConversionDiscount *string `json:"conversion_discount,omitempty"`
	// Is this an MFN (Most Favored Nations) flavored Convertible Note?
	ConversionMfn *bool `json:"conversion_mfn,omitempty"`
	// What is the valuation cap (if applicable)?
	ConversionValuationCap *types.Monetary `json:"conversion_valuation_cap,omitempty"`
	// How many days are there is a given period for calculation purposes?
	DayCountConvention enums.DayCountType `json:"day_count_convention"`
	// For cash proceeds calculation during a liquidity event.
	ExitMultiple *types.Ratio `json:"exit_multiple,omitempty"`
	// What is the period over which interest is calculated?
	InterestAccrualPeriod enums.AccrualPeriodType `json:"interest_accrual_period"`
	// How is interest paid out (if at applicable)
	InterestPayout enums.InterestPayoutType `json:"interest_payout"`
	// Interest rate(s) of the convertible (if applicable)
	InterestRates []types.InterestRate `json:"interest_rates"`
}
