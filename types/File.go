package types

// Type representation of a file
type File struct {
	// Path to the file within the OCF container
	Filepath string `json:"filepath"`
	// MD5 file checksum
	Md5 string `json:"md5"`
}
