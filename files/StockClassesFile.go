package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
)

type StockClassesFile struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// List of OCF stock class objects
	Items []objects.StockClass `json:"items"`
}
