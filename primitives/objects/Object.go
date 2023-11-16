package objects

import "github.com/tenkeylabs/go-ocf/enums"

// Abstract object to be extended by all other objects
type Object struct {
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Identifier for the object
	ID string `json:"id"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
}
