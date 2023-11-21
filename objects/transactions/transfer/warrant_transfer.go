package transfer

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/transfer"
)

// Object describing a transfer or secondary sale of a warrant security
type WarrantTransfer struct {
	transfer.Transfer
	transactions.Transaction
	transactions.SecurityTransaction
	objects.Object

	// Quantity of non-monetary security units transferred
	Quantity string `json:"quantity"`
}
