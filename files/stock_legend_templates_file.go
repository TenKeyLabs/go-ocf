package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list of stock legend templates
type StockLegendTemplatesFile struct {
	files.File

	// List of OCF stock legend template objects
	Items []objects.StockLegendTemplate `json:"items"`
}
