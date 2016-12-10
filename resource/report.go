package resource

import (
	. "github.com/talbright/go-desk/types"
	"time"
)

type InsightsV3Report struct {
	Request *InsightsV3Query `json:"request"`
	Header  []string         `json:"header"`
	RawData [][]interface{}  `json:"data"`
	Data    []*InsightsV3ReportItem
}

type InsightsV3ReportItem struct {
	Dimension1Value interface{}
	Dimension2Value interface{}
	WindowTime      Timestamp
	Data            map[string]interface{}
}

func (r *InsightsV3Report) Unravel() {
	for _, d := range r.RawData {
		// initialize the map
		item := &InsightsV3ReportItem{
			Data: make(map[string]interface{}),
		}

		// check that the window time is given
		if d[0] != nil {
			time, err := time.Parse("2006-01-02 15:04:05", d[0].(string))

			if err == nil {
				// add the window time as a timestamp type
				item.WindowTime = Timestamp{time}
			}
		}

		// dimension 1 value is always the second index
		item.Dimension1Value = d[1]
		// dimension 2 value is always the third index
		item.Dimension2Value = d[2]

		// need to grab a slice so the headers and data fields line up
		headers := r.Header[3:]
		// loop through the rest of the indexes and save the data to the items data field
		for j, val := range d[3:] {
			// use the headers value as the key in the map
			item.Data[headers[j]] = val
		}

		r.Data = append(r.Data, item)
	}
}
