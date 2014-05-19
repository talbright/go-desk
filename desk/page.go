package desk

import (
	"encoding/json"
)

type EntryCollection struct {
	RawEntries *json.RawMessage `json:"entries,omitempty"`
	Entries    []interface{}
}

type Page struct {
	PageNumber   *int                    `json:"page,omitempty"`
	TotalEntries *int                    `json:"total_entries,omitempty"`
	Embedded     *EntryCollection        `json:"_embedded,omitempty"`
  LinkCollection
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
