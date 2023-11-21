package conversiontriggers

import "github.com/tenkeylabs/go-ocf/primitives/types/conversiontriggers"

// Use this where no structured data is available regarding what triggers the conversion of
// a given security.
type UnspecifiedConversionTrigger struct {
	conversiontriggers.ConversionTrigger
}
