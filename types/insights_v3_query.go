package types

import "errors"

var (
	invalidWindowSize = errors.New("Invalid window size")
)

type InsightsV3Query struct {
	Fields           []string        `json:"fields,omitempty"`
	Dimension1       string          `json:"dimension1,omitempty"`
	Dimension1Values interface{}     `json:"dimension1_values,omitempty"`
	Dimension2       string          `json:"dimension2,omitempty"`
	Dimension2Values interface{}     `json:"dimension2_values,omitempty"`
	Time             *InsightsV3Time `json:"time"`
}

func NewInsightsV3Query() *InsightsV3Query {
	return &InsightsV3Query{
		Time: &InsightsV3Time{},
	}
}

type InsightsV3Time struct {
	Min        Timestamp `json:"min"`
	Max        Timestamp `json:"max"`
	WindowSize string    `json:"window_size"`
}

func (t *InsightsV3Time) SetWindowSize(size string) error {
	switch size {
	case "none", "hour", "day", "week", "month":
		t.WindowSize = size
		return nil
	}

	return invalidWindowSize
}
