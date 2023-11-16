package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
)

// JSON containing file type identifier and list of stock legend templates
//
// Abstract file to be extended by all other files
type StockLegendTemplatesFile struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// List of OCF stock legend template objects
	Items []objects.StockLegendTemplate `json:"items"`
}
