package transfer

import (
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions/transfer"
)

// Object describing a transfer of equity compensation
type EquityCompensationTransfer struct {
	transfer.Transfer
	transactions.SecurityTransaction

	// Quantity of non-monetary security units transferred
	Quantity string `json:"quantity"`
}
