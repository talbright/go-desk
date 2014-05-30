package desk

import (
  "github.com/lann/builder"
  "time"
)

type Message struct {
	Direction        *string    `json:"direction,omitempty"`
	Body             *string    `json:"body,omitempty"`
	BodyText         *string    `json:"body_text,omitempty"`
	BodyHtml         *string    `json:"body_html,omitempty"`
	Headers          *string    `json:"headers,omitempty"`
	HeadersRaw       *string    `json:"headers_raw,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Subject          *string    `json:"subject,omitempty"`
	To               *string    `json:"to,omitempty"`
	From             *string    `json:"from,omitempty"`
	Cc               *string    `json:"cc,omitempty"`
	Bcc              *string    `json:"bcc,omitempty"`
	ClientType       *string    `json:"client_type,omitempty"`
	FromFacebookName *string    `json:"from_facebook_name,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp `json:"updated_at,omitempty"`
	LinkCollection
}

func (c Message) String() string {
	return Stringify(c)
}

type messageBuilder builder.Builder

func (b messageBuilder) SetString(field string,value string) messageBuilder {
  return builder.Set(b, field, &value).(messageBuilder)
}

func (b messageBuilder) SetInt(field string,value int) messageBuilder {
  return builder.Set(b, field, &value).(messageBuilder)
}

func (b messageBuilder) SetTimestamp(field string,value Timestamp) messageBuilder {
  return builder.Set(b, field, &value).(messageBuilder)
}

func (b messageBuilder) SetTimestampNow(field string) messageBuilder {
  timet:=Timestamp{time.Now()}
  return builder.Set(b, field, &timet).(messageBuilder)
}

func (b messageBuilder) Build() Message {
  return builder.GetStruct(b).(Message)
}

var MessageBuilder = builder.Register(messageBuilder{}, Message{}).(messageBuilder)


