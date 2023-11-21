package conversionrights

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tenkeylabs/go-ocf/enums"
)

type ConversionRights struct {
	Type enums.ConversionRightType `json:"type,omitempty"`

	ConvertibleConversionRight *ConvertibleConversionRight `json:"convertible_conversion_right,omitempty"`
	StockClassConversionRight  *StockClassConversionRight  `json:"stock_class_conversion_right,omitempty"`
	WarrantConversionRight     *WarrantConversionRight     `json:"warrant_conversion_right,omitempty"`
}

func (c *ConversionRights) UnmarshalJSON(data []byte) error {
	aliasConversionRight := &struct {
		Type enums.ConversionRightType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &aliasConversionRight); err != nil {
		return err
	}

	c.Type = aliasConversionRight.Type

	var err error
	switch aliasConversionRight.Type {
	case enums.ConvertibleConversionRight:
		var convertibleConversionRight ConvertibleConversionRight
		err = json.Unmarshal(data, &convertibleConversionRight)
		c.ConvertibleConversionRight = &convertibleConversionRight
	case enums.StockClassConversionRight:
		var stockClassConversionRight StockClassConversionRight
		err = json.Unmarshal(data, &stockClassConversionRight)
		c.StockClassConversionRight = &stockClassConversionRight
	case enums.WarrantConversionRight:
		var warrantConversionRight WarrantConversionRight
		err = json.Unmarshal(data, &warrantConversionRight)
		c.WarrantConversionRight = &warrantConversionRight
	}

	return err
}

func (c ConversionRights) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case enums.ConvertibleConversionRight:
		return json.Marshal(c.ConvertibleConversionRight)
	case enums.StockClassConversionRight:
		return json.Marshal(c.StockClassConversionRight)
	case enums.WarrantConversionRight:
		return json.Marshal(c.WarrantConversionRight)
	}
	
	return nil, errors.New(fmt.Sprintf("unknown conversion right type: %s", c.Type))
}
