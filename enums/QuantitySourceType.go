package enums

// If quantity is provided, use this to specify where the number came from - e.g. was it a
// fixed value from the instrument (`INSTRUMENT_FIXED`), a human estimate
// (`HUMAN_ESTIMATED`), etc. If quantity is provided and this field is not, this is assumed
// to be `UNSPECIFIED`
//
// Enumeration of quantity source types describing where a quantity value came from
type QuantitySourceType string

const (
	QuantitySourceHumanEstimated   QuantitySourceType = "HUMAN_ESTIMATED"
	QuantitySourceInstrumentFixed  QuantitySourceType = "INSTRUMENT_FIXED"
	QuantitySourceInstrumentMax    QuantitySourceType = "INSTRUMENT_MAX"
	QuantitySourceInstrumentMin    QuantitySourceType = "INSTRUMENT_MIN"
	QuantitySourceMachineEstimated QuantitySourceType = "MACHINE_ESTIMATED"
	QuantitySourceUnspecified      QuantitySourceType = "UNSPECIFIED"
)
