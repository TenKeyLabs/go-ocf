package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list of stock plans
type StockPlansFile struct {
	files.File

	// List of OCF stock plan objects
	Items []objects.StockPlan `json:"items"`
}
