package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
)

// JSON containing file type identifier and list of valuations
//
// Abstract file to be extended by all other files
type ValuationsFile struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// List of OCF valuation objects
	Items []objects.Valuation `json:"items"`
}
