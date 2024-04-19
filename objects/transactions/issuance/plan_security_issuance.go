package issuance

// A Plan Security Issuance is a transaction to issue plan securities (it's a compatibility
// wrapper for Equity Compensation Issuances)
//
// Object describing securities issuance transaction by the issuer and held by a stakeholder
// as a form of compensation (as noted elsewhere, RSAs are not included here intentionally
// and should be modelled using Stock Issuances).
type PlanSecurityIssuance = EquityCompensationIssuance
