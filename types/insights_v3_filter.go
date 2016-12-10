//
// Copyright (c) 2015 Highwinds Network Group, inc.
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential.
//
// @author    Scot Wells <scot.wells@highwinds.com>
// @copyright 2015 Highwinds Network Group, inc.
//
package types

import (
	"encoding/json"
)

const (
	TYPE_INCLUDE = "Include"
	TYPE_EXCLUDE = "Exclude"

	// there's more fields available but they're currently undocumented...
	FIELD_CUSTOM_FIELDS = "Custom Fields"
	FIELD_LABELS        = "Labels"
)

type InsightsV3Filter struct {
	Type    string
	Field   string
	Filters []map[string]interface{}
}

func NewInsightsV3Filter() *InsightsV3Filter {
	return &InsightsV3Filter{
		Filters: []map[string]interface{}{},
	}
}

func (f *InsightsV3Filter) AddFilter(filter map[string]interface{}) {
	f.Filters = append(f.Filters, filter)
}

func (f *InsightsV3Filter) MarshalJSON() ([]byte, error) {
	data := [6]interface{}{}

	data[0] = f.Type
	data[1] = f.Field
	// this field is currently undocumented by Desk...
	// had to look at the request sent out by the Business Insights
	// dashboard just to get this to work properly
	data[2] = "Is"
	// same here...
	data[3] = ""
	data[4] = false
	data[5] = f.Filters

	j, err := json.Marshal(data)

	return j, err
}

type InsightsV3FilterGroup struct {
	Filters []*InsightsV3Filter
}

func NewInsightsV3FilterGroup() *InsightsV3FilterGroup {
	return &InsightsV3FilterGroup{}
}

func (g *InsightsV3FilterGroup) Add(filter *InsightsV3Filter) {
	g.Filters = append(g.Filters, filter)
}

func (f *InsightsV3FilterGroup) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(f.Filters)
	return j, err
}

// implements the Unmarshaler interface so when trying to parse the response it can
// be turned back into a filter struct
func (f *InsightsV3FilterGroup) UnmarshalJSON(d []byte) error {
	data := [][]interface{}{}

	err := json.Unmarshal(d, &data)

	if err != nil {
		return err
	}

	for _, f := range data {
		filter := NewInsightsV3Filter()

		filter.Type = f[0].(string)
		filter.Field = f[1].(string)

		for _, field := range f[5].([]interface{}) {
			filter.AddFilter(field.(map[string]interface{}))
		}
	}

	return nil
}
