package desk

import (
  "github.com/lann/builder"
  "time"
)

type Case struct {
	ID              *int                    `json:"id,omitempty"`
	ExternalID      *string                 `json:"external_id,omitempty"`
	Type            *string                 `json:"type,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	Subject         *string                 `json:"subject,omitempty"`
	Blurb           *string                 `json:"blurb,omitempty"`
	Language        *string                 `json:"language,omitempty"`
	Priority        *int                    `json:"priority,omitempty"`
	CustomFields    map[string]interface{}  `json:"custom_fields,omitempty"`
	LockedUntil     *Timestamp              `json:"locked_until",omitempty`
	CreatedAt       *Timestamp              `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp              `json:"updated_at,omitempty"`
	ReceivedAt      *Timestamp              `json:"received_at,omitempty"`
	ActiveAt        *Timestamp              `json:"active_at,omitempty"`
	OpenedAt        *Timestamp              `json:"opened_at,omitempty"`
	FirstOpenedAt   *Timestamp              `json:"first_opened_at,omitempty"`
	ResolvedAt      *Timestamp              `json:"resolved_at,omitempty"`
	FirstResolvedAt *Timestamp              `json:"first_resolved_at,omitempty"`
  Message         *Message                `json:"message,omitempty"`
	LinkCollection
}

type caseBuilder builder.Builder

func (c Case) String() string {
	return Stringify(c)
}

func (b caseBuilder) SetString(field string,value string) caseBuilder {
  return builder.Set(b, field, &value).(caseBuilder)
}

func (b caseBuilder) SetInt(field string,value int) caseBuilder {
  return builder.Set(b, field, &value).(caseBuilder)
}

func (b caseBuilder) SetTimestamp(field string,value Timestamp) caseBuilder {
  return builder.Set(b, field, &value).(caseBuilder)
}

func (b caseBuilder) SetTimestampNow(field string) caseBuilder {
  timet:=Timestamp{time.Now()}
  return builder.Set(b, field, &timet).(caseBuilder)
}

func (b caseBuilder) SetMessage(value Message) caseBuilder {
  return builder.Set(b, "Message", &value).(caseBuilder)
}

func (b caseBuilder) SetCustomField(name string,value interface{}) caseBuilder {
  var fields map[string]interface{}
  caze:=builder.GetStruct(b).(Case)
  if caze.CustomFields==nil {
    fields=make(map[string]interface{})
  } else {
    fields=caze.CustomFields
  }
  fields[name]=value
  return builder.Set(b, "CustomFields", fields).(caseBuilder)
}

func (b caseBuilder) SetLinkCollection(value LinkCollection) caseBuilder {
  return builder.Set(b, "LinkCollection", value).(caseBuilder)
}

func (b caseBuilder) Build() Case {
  return builder.GetStruct(b).(Case)
}

var CaseBuilder = builder.Register(caseBuilder{}, Case{}).(caseBuilder)

