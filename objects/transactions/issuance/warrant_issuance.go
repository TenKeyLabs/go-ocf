package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

// Object describing warrant issuance transaction by the issuer and held by a stakeholder
type WarrantIssuance struct {
	issuance.Issuance
	transactions.SecurityTransaction

	// The exercise price of the warrant
	ExercisePrice *types.Monetary `json:"exercise_price,omitempty"`
	// In event the Warrant can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ExerciseTriggers []conversiontriggers.ConversionTrigger `json:"exercise_triggers"`
	// Actual purchase price of the warrant (sum up purported value of all consideration,
	// including in-kind)
	PurchasePrice types.Monetary `json:"purchase_price"`
	// Quantity of shares the warrant is exercisable for
	Quantity *string `json:"quantity,omitempty"`
	// If quantity is provided, use this to specify where the number came from - e.g. was it a
	// fixed value from the instrument (`INSTRUMENT_FIXED`), a human estimate
	// (`HUMAN_ESTIMATED`), etc. If quantity is provided and this field is not, this is assumed
	// to be `UNSPECIFIED`
	QuantitySource *enums.QuantitySourceType `json:"quantity_source,omitempty"`
	// Identifier of the VestingTerms to which this security is subject. If not present,
	// security is fully vested on issuance.
	VestingTermsID *string `json:"vesting_terms_id,omitempty"`
	// What is expiration date of the warrant (if applicable)
	WarrantExpirationDate *string `json:"warrant_expiration_date,omitempty"`
}
