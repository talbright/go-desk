package desk

type Case struct {
	ID              *int                    `json:"id,omitempty"`
	ExternalID      *string                 `json:"external_id,omitempty"`
	Links           *map[string]interface{} `json:"_links,omitempty"`
	Type            *string                 `json:"type,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	Subject         *string                 `json:"subject,omitempty"`
	Blurb           *string                 `json:"blurb,omitempty"`
	Language        *string                 `json:"language,omitempty"`
	Priority        *int                    `json:"priority,omitempty"`
	CustomFields    *map[string]interface{} `json:"custom_fields,omitempty"`
	LockedUntil     *Timestamp              `json:"locked_until",omitempty`
	CreatedAt       *Timestamp              `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp              `json:"updated_at,omitempty"`
	ReceivedAt      *Timestamp              `json:"received_at,omitempty"`
	ActiveAt        *Timestamp              `json:"active_at,omitempty"`
	OpenedAt        *Timestamp              `json:"opened_at,omitempty"`
	FirstOpenedAt   *Timestamp              `json:"first_opened_at,omitempty"`
	ResolvedAt      *Timestamp              `json:"resolved_at,omitempty"`
	FirstResolvedAt *Timestamp              `json:"first_resolved_at,omitempty"`
}

func (c Case) String() string {
	return Stringify(c)
}
