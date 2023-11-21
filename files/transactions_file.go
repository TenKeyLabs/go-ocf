package files

import (
	"encoding/json"

	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list transactions
type TransactionsFile struct {
	files.File

	// List of OCF transaction objects
	Items []map[string]*json.RawMessage `json:"items"` //TODO: Union types
}
