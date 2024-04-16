package enums

// The calendar day of a month to award vesting.
//
// Enumeration representing a vesting "day of month". Since not all months have 29, 30, or
// 31 days, this enum requires those values to also specify an overflow behavior.
// - `01` - `28` : Day 1, 2... 28 of the month; e.g. `03` means vesting occurs on the 3rd of
// the month.
// - `29_OR_LAST_DAY_OF_MONTH` - `31_OR_LAST_DAY_OF_MONTH` : Day 29, 30, or 31 of the month,
// defaulting to the last day of the month for shorter months; e.g.
// `31_OR_LAST_DAY_OF_MONTH` means monthly vesting occurs on Jan 31, Feb 28/29, Mar 31, Apr
// 30, etc.
// - `VESTING_START_DAY_OR_LAST_DAY_OF_MONTH` vests on the same day of month as the day of
// the `VESTING_START` condition; e.g. if vesting commences on Jan 15 then any vesting will
// accrue on the 15th of future vesting months. If vesting commencement occurs on days
// 29-31, this has the same behavior as the other `*_LAST_DAY_OF_MONTH` values.
type VestingDayOfMonth string

const (
	The01                           VestingDayOfMonth = "01"
	The02                           VestingDayOfMonth = "02"
	The03                           VestingDayOfMonth = "03"
	The04                           VestingDayOfMonth = "04"
	The05                           VestingDayOfMonth = "05"
	The06                           VestingDayOfMonth = "06"
	The07                           VestingDayOfMonth = "07"
	The08                           VestingDayOfMonth = "08"
	The09                           VestingDayOfMonth = "09"
	The10                           VestingDayOfMonth = "10"
	The11                           VestingDayOfMonth = "11"
	The12                           VestingDayOfMonth = "12"
	The13                           VestingDayOfMonth = "13"
	The14                           VestingDayOfMonth = "14"
	The15                           VestingDayOfMonth = "15"
	The16                           VestingDayOfMonth = "16"
	The17                           VestingDayOfMonth = "17"
	The18                           VestingDayOfMonth = "18"
	The19                           VestingDayOfMonth = "19"
	The20                           VestingDayOfMonth = "20"
	The21                           VestingDayOfMonth = "21"
	The22                           VestingDayOfMonth = "22"
	The23                           VestingDayOfMonth = "23"
	The24                           VestingDayOfMonth = "24"
	The25                           VestingDayOfMonth = "25"
	The26                           VestingDayOfMonth = "26"
	The27                           VestingDayOfMonth = "27"
	The28                           VestingDayOfMonth = "28"
	The29_OrLastDayOfMonth          VestingDayOfMonth = "29_OR_LAST_DAY_OF_MONTH"
	The30_OrLastDayOfMonth          VestingDayOfMonth = "30_OR_LAST_DAY_OF_MONTH"
	The31_OrLastDayOfMonth          VestingDayOfMonth = "31_OR_LAST_DAY_OF_MONTH"
	VestingStartDayOrLastDayOfMonth VestingDayOfMonth = "VESTING_START_DAY_OR_LAST_DAY_OF_MONTH"
)
