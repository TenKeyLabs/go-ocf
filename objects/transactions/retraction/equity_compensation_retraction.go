package retraction

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/retraction"
)

// Object describing a retraction of equity compensation
type EquityCompensationRetraction struct {
	retraction.Retraction
	transactions.SecurityTransaction
}
