package types

import (
	"errors"
)

var (
	invalidWindowSize = errors.New("Invalid window size")
)

type InsightsV3Time struct {
	Min        Timestamp `json:"min"`
	Max        Timestamp `json:"max"`
	WindowSize string    `json:"window_size"`
}

func NewInsightsV3Time() *InsightsV3Time {
	return &InsightsV3Time{
		WindowSize: "none",
	}
}

func (t *InsightsV3Time) SetWindowSize(size string) error {
	switch size {
	case "none", "hour", "day", "week", "month":
		t.WindowSize = size
		return nil
	}

	return invalidWindowSize
}
