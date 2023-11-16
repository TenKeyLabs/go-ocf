package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/vesting"
)

// Object describing the terms under which a security vests
//
// Abstract object to be extended by all other objects
type VestingTerms struct {
	// Allocation/rounding type for the vesting schedule
	AllocationType enums.AllocationType `json:"allocation_type"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// Detailed description of the vesting schedule
	Description string `json:"description"`
	// Identifier for the object
	ID string `json:"id"`
	// Concise name for the vesting schedule
	Name string `json:"name"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Conditions and triggers that describe the graph of vesting schedules and events
	VestingConditions []vesting.VestingCondition `json:"vesting_conditions"`
}
