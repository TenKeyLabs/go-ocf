package conversionmechanisms

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tenkeylabs/go-ocf/enums"
)

type ConversionMechanism struct {
	Type enums.ConversionMechanismType `json:"type"`

	CustomConversionMechanism                *CustomConversionMechanism                `json:"custom_conversion_mechanism,omitempty"`
	FixedAmountConversionMechanism           *FixedAmountConversionMechanism           `json:"fixed_amount_conversion_mechanism,omitempty"`
	NoteConversionMechanism                  *NoteConversionMechanism                  `json:"note_conversion_mechanism,omitempty"`
	PercentCapitalizationConversionMechanism *PercentCapitalizationConversionMechanism `json:"percent_capitalization_conversion_mechanism,omitempty"`
	RatioConversionMechanism                 *RatioConversionMechanism                 `json:"ratio_conversion_mechanism,omitempty"`
	SAFEConversionMechanism                  *SAFEConversionMechanism                  `json:"safe_conversion_mechanism,omitempty"`
	SharePriceBasedConversionMechanism       *SharePriceBasedConversionMechanism       `json:"share_price_based_conversion_mechanism,omitempty"`
	ValuationBasedConversionMechanism        *ValuationBasedConversionMechanism        `json:"valuation_based_conversion_mechanism,omitempty"`
}

func (c *ConversionMechanism) UnmarshalJSON(data []byte) error {
	aliasConversionMechanism := &struct {
		Type enums.ConversionMechanismType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &aliasConversionMechanism); err != nil {
		return err
	}

	c.Type = aliasConversionMechanism.Type

	var err error
	switch aliasConversionMechanism.Type {
	case enums.CustomConversion:
		var customConversionMechanism CustomConversionMechanism
		err = json.Unmarshal(data, &customConversionMechanism)
		c.CustomConversionMechanism = &customConversionMechanism
	case enums.FixedAmountConversion:
		var fixedAmountConversionMechanism FixedAmountConversionMechanism
		err = json.Unmarshal(data, &fixedAmountConversionMechanism)
		c.FixedAmountConversionMechanism = &fixedAmountConversionMechanism
	case enums.ConvertibleNoteConversion:
		var noteConversionMechanism NoteConversionMechanism
		err = json.Unmarshal(data, &noteConversionMechanism)
		c.NoteConversionMechanism = &noteConversionMechanism
	case enums.FixedPercentOfCapitalizationConversion:
		var percentCapitalizationConversionMechanism PercentCapitalizationConversionMechanism
		err = json.Unmarshal(data, &percentCapitalizationConversionMechanism)
		c.PercentCapitalizationConversionMechanism = &percentCapitalizationConversionMechanism
	case enums.RatioConversion:
		var ratioConversionMechanism RatioConversionMechanism
		err = json.Unmarshal(data, &ratioConversionMechanism)
		c.RatioConversionMechanism = &ratioConversionMechanism
	case enums.SAFEConversion:
		var safeConversionMechanism SAFEConversionMechanism
		err = json.Unmarshal(data, &safeConversionMechanism)
		c.SAFEConversionMechanism = &safeConversionMechanism
	case enums.PpsBasedConversion:
		var sharePriceBasedConversionMechanism SharePriceBasedConversionMechanism
		err = json.Unmarshal(data, &sharePriceBasedConversionMechanism)
		c.SharePriceBasedConversionMechanism = &sharePriceBasedConversionMechanism
	case enums.ValuationBasedConversion:
		var valuationBasedConversionMechanism ValuationBasedConversionMechanism
		err = json.Unmarshal(data, &valuationBasedConversionMechanism)
		c.ValuationBasedConversionMechanism = &valuationBasedConversionMechanism
	}

	return err
}

func (c *ConversionMechanism) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case enums.CustomConversion:
		return json.Marshal(c.CustomConversionMechanism)
	case enums.FixedAmountConversion:
		return json.Marshal(c.FixedAmountConversionMechanism)
	case enums.ConvertibleNoteConversion:
		return json.Marshal(c.NoteConversionMechanism)
	case enums.FixedPercentOfCapitalizationConversion:
		return json.Marshal(c.PercentCapitalizationConversionMechanism)
	case enums.RatioConversion:
		return json.Marshal(c.RatioConversionMechanism)
	case enums.SAFEConversion:
		return json.Marshal(c.SAFEConversionMechanism)
	case enums.PpsBasedConversion:
		return json.Marshal(c.SharePriceBasedConversionMechanism)
	case enums.ValuationBasedConversion:
		return json.Marshal(c.ValuationBasedConversionMechanism)
	}

	return nil, errors.New(fmt.Sprintf("unknown conversion mechanism type: %v", c.Type))
}
