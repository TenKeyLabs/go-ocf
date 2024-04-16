package conversiontriggers

import "github.com/tenkeylabs/go-ocf/primitives/types/conversiontriggers"

// Type representation of elective trigger valid at will (so long as instrument is valid and
// outstanding).
type ElectiveConversionAtWillTrigger struct {
	conversiontriggers.ConversionTrigger
}
