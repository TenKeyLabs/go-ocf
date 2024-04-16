package files

import (
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/primitives/files"
)

// JSON containing file type identifier and list of stakeholders
type StakeholdersFile struct {
	files.File
	
	// List of OCF stakeholder objects
	Items []objects.Stakeholder `json:"items"`
}
