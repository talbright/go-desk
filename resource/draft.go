package resource

import(
	. "github.com/talbright/go-desk/types"
)

type Draft struct {
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
	Type             *string    `json:"type,omitempty"`
	Cc               *string    `json:"cc,omitempty"`
	Bcc              *string    `json:"bcc,omitempty"`
	ClientType       *string    `json:"client_type,omitempty"`
	FromFacebookName *string    `json:"from_facebook_name,omitempty"`
	PublicUrl        *string    `json:"public_url,omitempty"`
	IsBestAnswer     *string    `json:"is_best_answer,omitempty"`
	Rating           *float32   `json:"rating,omitempty"`
	RatingCount      *int       `json:"rating_count,omitempty"`
	RatingScore      *int       `json:"rating_score,omitempty"`
	EnteredAt        *Timestamp `json:"entered_at,omitempty"`
	HiddentAt        *Timestamp `json:"hidden_at,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp `json:"updated_at,omitempty"`
	Resource
}

func NewDraft() *Draft {
	d := &Draft{}
	d.InitializeResource(d)
	d.Singularize()
	return d
}

func (c Draft) String() string {
	return Stringify(c)
}
