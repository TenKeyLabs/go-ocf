package retraction

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/retraction"
)

// Object describing a retraction of a warrant security
type WarrantRetraction struct {
	retraction.Retraction
	transactions.SecurityTransaction
}
