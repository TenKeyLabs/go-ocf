package release

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/release"
)

// Object describing equity compensation security release transaction
type EquityCompensationRelease struct {
	release.Release
	transactions.SecurityTransaction
}
