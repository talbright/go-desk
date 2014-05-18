package desk

import (
	"encoding/json"
)

type Embedded struct {
	RawEntries *json.RawMessage `json:"entries,omitempty"`
	Entries    []interface{}
}

type Collection struct {
	Page         *int                    `json:"page,omitempty"`
	TotalEntries *int                    `json:"total_entries,omitempty"`
	Links        *map[string]interface{} `json:"_links,omitempty"`
	Embed        *Embedded               `json:"_embedded,omitempty"`
}

func (c Collection) String() string {
	return Stringify(c)
}

//TODO still a work in progress
//func (e *Embedded) UnmarshalRawEntries(entries interface{}) (error) {
//
// err := json.Unmarshal(*e.RawEntries,entries)
// if err != nil {
//	  return err
//	}
//  log.Printf("unmarshall raw: %v",entries)
//	e.Entries = make([]interface{},len(&entries))
//  return nil
//}
