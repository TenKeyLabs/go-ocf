package enums

// Allocation/rounding type for the vesting schedule
//
// Enumeration of allocation types for vesting terms. Using an example of 18 shares split
// across 4 tranches, each allocation type results in a different schedule as follows:
// 1.  Cumulative Rounding (5 - 4 - 5 - 4)
// 2.  Cumulative Round Down (4 - 5 - 4 - 5)
// 3.  Front Loaded (5 - 5 - 4 - 4)
// 4.  Back Loaded (4 - 4 - 5 - 5)
// 5.  Front Loaded to Single Tranche (6 - 4 - 4 - 4)
// 6.  Back Loaded to Single Tranche (4 - 4 - 4 - 6)
// 7.  Fractional (4.5 - 4.5 - 4.5 - 4.5)
type AllocationType string

const (
	BackLoaded                 AllocationType = "BACK_LOADED"
	BackLoadedToSingleTranche  AllocationType = "BACK_LOADED_TO_SINGLE_TRANCHE"
	CumulativeRoundDown        AllocationType = "CUMULATIVE_ROUND_DOWN"
	CumulativeRounding         AllocationType = "CUMULATIVE_ROUNDING"
	Fractional                 AllocationType = "FRACTIONAL"
	FrontLoaded                AllocationType = "FRONT_LOADED"
	FrontLoadedToSingleTranche AllocationType = "FRONT_LOADED_TO_SINGLE_TRANCHE"
)
