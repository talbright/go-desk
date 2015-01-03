package desk

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
	Hal
}

func (c Message) String() string {
	return Stringify(c)
}
