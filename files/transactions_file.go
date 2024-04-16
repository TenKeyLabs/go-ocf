package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list transactions
type TransactionsFile struct {
	files.File

	// List of OCF transaction objects
	Items []objects.Transaction `json:"items"`
}
