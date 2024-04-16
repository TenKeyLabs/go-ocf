package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/types/vesting"
)

// Object describing the terms under which a security vests
type VestingTerms struct {
	objects.Object

	// Allocation/rounding type for the vesting schedule
	AllocationType enums.AllocationType `json:"allocation_type"`
	// Detailed description of the vesting schedule
	Description string `json:"description"`
	// Concise name for the vesting schedule
	Name string `json:"name"`
	// Conditions and triggers that describe the graph of vesting schedules and events
	VestingConditions []vesting.VestingCondition `json:"vesting_conditions"`
}
