package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

type StockClassesFile struct {
	files.File

	// List of OCF stock class objects
	Items []objects.StockClass `json:"items"`
}
