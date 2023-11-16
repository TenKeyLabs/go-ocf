package files

import "github.com/tenkeylabs/go-ocf/enums"

// Abstract file to be extended by all other files
type File struct {
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
}
