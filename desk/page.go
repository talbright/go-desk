package desk

import (
	"encoding/json"
)

// EntryCollection holds the raw json data for embedded resources.
// The RawEntries field is used internally and not something an
// API user would typically want to access.
//
// Once the EntryCollection has been unmarshaled these resources
// become available as typed GO objects in the Entries field,
// which contains the data you probably want to get at.
type EntryCollection struct {
	RawEntries *json.RawMessage `json:"entries,omitempty"`
	Entries    []interface{}
}

// Page represents a single page of results, typically from a search
// or list API method. A page has an embedded collection of resources
// which, contains the data you probably want to get at.
// See Desk API (http://dev.desk.com/API/using-the-api/#embedding)
type Page struct {
	PageNumber   *int                              `json:"page,omitempty"`
	TotalEntries *int                              `json:"total_entries,omitempty"`
	Embedded     *EntryCollection                  `json:"_embedded,omitempty"`
	Links        map[string]map[string]interface{} `json:"_links,omitempty"`
}

func (c Page) String() string {
	return Stringify(c)
}

//TODO create generalized method for unmarshalling raw entries by type
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
