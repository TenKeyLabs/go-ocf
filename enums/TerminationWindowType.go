package enums

// What cause of termination is this window for?
//
// Enumeration of termination window types
type TerminationWindowType string

const (
	InvoluntaryDeath      TerminationWindowType = "INVOLUNTARY_DEATH"
	InvoluntaryDisability TerminationWindowType = "INVOLUNTARY_DISABILITY"
	InvoluntaryOther      TerminationWindowType = "INVOLUNTARY_OTHER"
	InvoluntaryWithCause  TerminationWindowType = "INVOLUNTARY_WITH_CAUSE"
	VoluntaryGoodCause    TerminationWindowType = "VOLUNTARY_GOOD_CAUSE"
	VoluntaryOther        TerminationWindowType = "VOLUNTARY_OTHER"
	VoluntaryRetirement   TerminationWindowType = "VOLUNTARY_RETIREMENT"
)
