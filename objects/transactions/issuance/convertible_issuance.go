package issuance

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversiontriggers"
)

// Object describing convertible instrument issuance transaction by the issuer and held by a
// stakeholder
type ConvertibleIssuance struct {
	issuance.Issuance
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// In event the convertible can convert due to trigger events (e.g. Maturity, Next Qualified
	// Financing, Change of Control, at Election of Holder), what are the terms?
	ConversionTriggers []conversiontriggers.ConversionTrigger `json:"conversion_triggers"`
	// What kind of convertible instrument is this (of the supported, enumerated types)
	ConvertibleType enums.ConvertibleType `json:"convertible_type"`
	// Amount invested and outstanding on date of issuance of this convertible
	InvestmentAmount types.Monetary `json:"investment_amount"`
	// What pro-rata (if any) is the holder entitled to buy at the next round?
	ProRata *string `json:"pro_rata,omitempty"`
	// If different convertible instruments have seniorty over one another, use this value to
	// build a seniority stack, with 1 being highest seniority and equal seniority values
	// assumed to be equal priority
	Seniority int64 `json:"seniority"`
}
