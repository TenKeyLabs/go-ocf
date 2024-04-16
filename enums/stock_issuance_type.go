package enums

// Optional field to flag certain special types of issuances (like RSAs)
//
// Enumeration of issuance types where we want to draw attention to some unique aspect of a
// stock issuance (e.g. is it an RSA)
type StockIssuanceType string

const (
	FoundersStock StockIssuanceType = "FOUNDERS_STOCK"
	RSA           StockIssuanceType = "RSA"
)
