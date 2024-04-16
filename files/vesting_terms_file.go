package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list of vesting terms
type VestingTermsFile struct {
	files.File

	// List of OCF vesting terms objects
	Items []objects.VestingTerms `json:"items"`
}
