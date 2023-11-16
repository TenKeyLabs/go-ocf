package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
)

// JSON containing file type identifier and list transactions
//
// Abstract file to be extended by all other files
type TransactionsFile struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// List of OCF transaction objects
	Items []objects.Transaction `json:"items"` //TODO: Union types
}
