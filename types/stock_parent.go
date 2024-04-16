package types

import "github.com/tenkeylabs/go-ocf/enums"

// Type representation of the parent security of a given stock issuance (e.g. if a stock
// issuance came from a plan, such as an RSA, or if a stock came from a previous stock entry)
type StockParent struct {
	// Parent object's ID must be a valid ID pointing to an object of the type specified in
	// parent_object_type
	ParentObjectID string `json:"parent_object_id"`
	// Parent object type for this stock issuance (e.g. a stock plan or warrant)
	ParentObjectType enums.ParentSecurityType `json:"parent_object_type"`
}
