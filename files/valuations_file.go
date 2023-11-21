package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list of valuations
type ValuationsFile struct {
	files.File

	// List of OCF valuation objects
	Items []objects.Valuation `json:"items"`
}
