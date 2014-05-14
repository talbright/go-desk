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

func (e *Embedded) UnmarshalRawEntries() {
	//TODO not sure how to do this in Go, but it would be nice to
	//make the code snippet below generic enough, so that if we
	//pass in the type to this method, we can unmarshal accordingly
	// cases := new([]Case)
	// err = json.Unmarshal(*collection.Embed.RawEntries,&cases)
	// if err != nil {
	//   return nil,resp,err
	// }
	// collection.Embed.Entries = make([]interface{},len(*cases))
	// for i,v := range *cases {
	//   collection.Embed.Entries[i] = interface{}(v)
	// }
	// collection.Embed.RawEntries = nil
}
