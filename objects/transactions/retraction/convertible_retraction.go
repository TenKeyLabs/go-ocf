package retraction

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/retraction"
)

// Object describing a retraction of a convertible security
type ConvertibleRetraction struct {
	retraction.Retraction
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object
}