package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
)

// JSON containing file type identifier and list of vesting terms
//
// Abstract file to be extended by all other files
type VestingTermsFile struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// List of OCF vesting terms objects
	Items []objects.VestingTerms `json:"items"`
}
