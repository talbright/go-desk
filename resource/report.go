package resource

import (
	. "github.com/talbright/go-desk/types"
)


type InsightsV3Report struct {
	Request *InsightsV3Query `json:"request"`
	Header  []string         `json:"header"`
	Data    [][]interface{}  `json:"data"`
}

func (r *InsightsV3Report) GetData() []map[string] interface{} {
	data := []map[string] interface{}{}

	for i, d := range r.Data {
		// initialize the map
		data = append(data, map[string] interface{}{})

		// save all the values in the map using the header as the template
		for j, header := range r.Header {
			data[i][header] = d[j]
		}
	}

	return data
}
