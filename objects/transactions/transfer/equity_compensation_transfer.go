package transfer

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/transfer"
)

// Object describing a transfer of equity compensation
type EquityCompensationTransfer struct {
	transfer.Transfer
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Quantity of non-monetary security units transferred
	Quantity string `json:"quantity"`
}
