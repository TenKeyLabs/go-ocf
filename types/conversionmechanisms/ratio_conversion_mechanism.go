package conversionmechanisms

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/primitives/types/conversionmechanisms"
	"github.com/tenkeylabs/go-ocf/types"
)

// Sets forth inputs and conversion mechanism of a ratio conversion (primarily used to
// describe conversion from one stock class (e.g. Preferred) into another (e.g. Common)
type RatioConversionMechanism struct {
	conversionmechanisms.ConversionMechanism

	// What is the effective conversion price per share of this stock class?
	ConversionPrice types.Monetary `json:"conversion_price"`
	// One share of this stock class converts into this many target stock class shares
	Ratio types.Ratio `json:"ratio"`
	// How should fractional shares be rounded?
	RoundingType enums.RoundingType `json:"rounding_type"`
}
