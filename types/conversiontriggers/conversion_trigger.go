package conversiontriggers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tenkeylabs/go-ocf/enums"
)

type ConversionTrigger struct {
	Type enums.ConversionTriggerType `json:"type,omitempty"`

	AutomaticConversionOnConditionTrigger *AutomaticConversionOnConditionTrigger `json:"automatic_conversion_on_condition_trigger,omitempty"`
	AutomaticConversionOnDateTrigger      *AutomaticConversionOnDateTrigger      `json:"automatic_conversion_on_date_trigger,omitempty"`
	ElectiveConversionAtWillTrigger       *ElectiveConversionAtWillTrigger       `json:"elective_conversion_at_will_trigger,omitempty"`
	ElectiveConversionInDateRangeTrigger  *ElectiveConversionInDateRangeTrigger  `json:"elective_conversion_in_date_range_trigger,omitempty"`
	ElectiveConversionOnConditionTrigger  *ElectiveConversionOnConditionTrigger  `json:"elective_conversion_on_condition_trigger,omitempty"`
	UnspecifiedConversionTrigger          *UnspecifiedConversionTrigger          `json:"unspecified_conversion_trigger,omitempty"`
}

func (c *ConversionTrigger) UnmarshalJSON(data []byte) error {
	aliasConversionTrigger := &struct {
		Type enums.ConversionTriggerType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &aliasConversionTrigger); err != nil {
		return err
	}

	c.Type = aliasConversionTrigger.Type

	var err error
	switch aliasConversionTrigger.Type {
	case enums.AutomaticOnCondition:
		var automaticConversionOnConditionTrigger AutomaticConversionOnConditionTrigger
		err = json.Unmarshal(data, &automaticConversionOnConditionTrigger)
		c.AutomaticConversionOnConditionTrigger = &automaticConversionOnConditionTrigger
	case enums.AutomaticOnDate:
		var automaticConversionOnDateTrigger AutomaticConversionOnDateTrigger
		err = json.Unmarshal(data, &automaticConversionOnDateTrigger)
		c.AutomaticConversionOnDateTrigger = &automaticConversionOnDateTrigger
	case enums.ElectiveAtWill:
		var electiveConversionAtWillTrigger ElectiveConversionAtWillTrigger
		err = json.Unmarshal(data, &electiveConversionAtWillTrigger)
		c.ElectiveConversionAtWillTrigger = &electiveConversionAtWillTrigger
	case enums.ElectiveInRange:
		var electiveConversionInDateRangeTrigger ElectiveConversionInDateRangeTrigger
		err = json.Unmarshal(data, &electiveConversionInDateRangeTrigger)
		c.ElectiveConversionInDateRangeTrigger = &electiveConversionInDateRangeTrigger
	case enums.ElectiveOnCondition:
		var electiveConversionOnConditionTrigger ElectiveConversionOnConditionTrigger
		err = json.Unmarshal(data, &electiveConversionOnConditionTrigger)
		c.ElectiveConversionOnConditionTrigger = &electiveConversionOnConditionTrigger
	case enums.Unspecified:
		var unspecifiedConversionTrigger UnspecifiedConversionTrigger
		err = json.Unmarshal(data, &unspecifiedConversionTrigger)
		c.UnspecifiedConversionTrigger = &unspecifiedConversionTrigger
	}

	return err
}

func (c ConversionTrigger) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case enums.AutomaticOnCondition:
		return json.Marshal(c.AutomaticConversionOnConditionTrigger)
	case enums.AutomaticOnDate:
		return json.Marshal(c.AutomaticConversionOnDateTrigger)
	case enums.ElectiveAtWill:
		return json.Marshal(c.ElectiveConversionAtWillTrigger)
	case enums.ElectiveInRange:
		return json.Marshal(c.ElectiveConversionInDateRangeTrigger)
	case enums.ElectiveOnCondition:
		return json.Marshal(c.ElectiveConversionOnConditionTrigger)
	case enums.Unspecified:
		return json.Marshal(c.UnspecifiedConversionTrigger)
	}

	return nil, errors.New(fmt.Sprintf("unknown conversion trigger type: %s", c.Type))
}
