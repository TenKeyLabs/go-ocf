package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/vesting"
)

// Describes a vesting condition satisfied at the security's vesting commencement date
type VestingStartTrigger struct {
	vesting.VestingConditionTrigger
}
