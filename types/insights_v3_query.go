package types

import (
	"encoding/json"
)

// allows easy creation of an insights query for the Desk Reporting API
type InsightsV3Query struct {
	Fields           []string                 `json:"fields,omitempty"`
	Dimension1       string                   `json:"dimension1,omitempty"`
	Dimension1Values interface{}              `json:"dimension1_values,omitempty"`
	Dimension2       string                   `json:"dimension2,omitempty"`
	Dimension2Values interface{}              `json:"dimension2_values,omitempty"`
	Time             *InsightsV3Time          `json:"time"`
	Filters          []*InsightsV3FilterGroup `json:"filters,omitempty"`
	// keeping a map of all the fields allows us quickly check if the
	// field was already added or not, otherwise we'd have to loop through
	// the entire array of fields to make sure it is unique
	fieldMap map[string]bool
}

// create a new insights query
func NewInsightsV3Query() *InsightsV3Query {
	return &InsightsV3Query{
		Time:     NewInsightsV3Time(),
		fieldMap: map[string]bool{},
		Filters:  []*InsightsV3FilterGroup{},
		Fields:   []string{},
	}
}

// adds a single field to the query
func (q *InsightsV3Query) AddField(field string) *InsightsV3Query {
	if _, ok := q.fieldMap[field]; !ok {
		// append the field to the query
		q.Fields = append(q.Fields, field)
		// marks that the field has been added so we know items are unique
		q.fieldMap[field] = true
	}

	return q
}

// add multiple fields to the query
func (q *InsightsV3Query) AddFields(fields []string) *InsightsV3Query {
	for _, field := range fields {
		// pass each field on to the add field function
		q.AddField(field)
	}

	return q
}

func (q *InsightsV3Query) SetReportTime(time *InsightsV3Time) *InsightsV3Query {
	q.Time = time

	return q
}

func (q *InsightsV3Query) SetDimension1(dimension string) *InsightsV3Query {
	q.Dimension1 = dimension

	return q
}

func (q *InsightsV3Query) SetDimension1Values(value interface{}) *InsightsV3Query {
	q.Dimension1Values = value

	return q
}

func (q *InsightsV3Query) SetDimension2(dimension string) *InsightsV3Query {
	q.Dimension2 = dimension

	return q
}

func (q *InsightsV3Query) SetDimension2Values(value interface{}) *InsightsV3Query {
	q.Dimension2Values = value

	return q
}

// add a filter to the query
func (q *InsightsV3Query) AddFilters(filters *InsightsV3FilterGroup) *InsightsV3Query {
	q.Filters = append(q.Filters, filters)

	return q
}

func (t *InsightsV3Query) String() string {
	j, err := json.Marshal(t)

	if err != nil {
		return ""
	}

	return string(j)
}
