package conversionrights

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types/conversionmechanisms"
)

// Type representation of a conversion right from one Stock Class into another Stock Class
//
// Abstract type representation of a conversion right from a non-plan security into another
// non-plan security
type StockClassConversionRights struct {
	// What conversion mechanism applies to calculate the number of resulting securities?
	ConversionMechanism conversionmechanisms.ConversionMechanism `json:"conversion_mechanism"`
	// Is this stock class potentially convertible into a future, as-yet undetermined stock
	// class (e.g. Founder Preferred)
	ConvertsToFutureRound *bool `json:"converts_to_future_round,omitempty"`
	// The identifier of the existing, known stock class this stock class can convert into
	ConvertsToStockClassID *string `json:"converts_to_stock_class_id,omitempty"`
	// What kind of conversion right is this?
	Type *enums.ConversionRightType `json:"type,omitempty"`
}
