package vesting

import (
	"github.com/tenkeylabs/go-ocf/primitives/types/vesting"
)

// Describes a period of time expressed in days (e.g. 365 days) for use in Vesting Terms
type VestingPeriodInDays struct {
	vesting.VestingPeriod
}
