package desk

type Customer struct {
	ID                  *int          `json:"id,omitempty"`
	ExternalID          *string       `json:"external_id,omitempty"`
	FirstName           *string       `json:"first_name,omitempty"` 
	LastName            *string       `json:"last_name,omitempty"` 
	Company             *string       `json:"company,omitempty"` 
	Title               *string       `json:"title,omitempty"` 
	Avatar              *string       `json:"avatar,omitempty"` 
	Background          *string       `json:"background,omitempty"` 
	Language            *string       `json:"language,omitempty"` 
	LockedUntil         *Timestamp    `json:"locked_until,omitempty"`
	CreatedAt           *Timestamp    `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp    `json:"updated_at,omitempty"`
	CustomFields        map[string]interface{}    `json:"custom_fields,omitempty"` 
	Emails              []interface{}  `json:"emails,omitempty"` 
	PhoneNumbers        []interface{}  `json:"phone_numbers,omitempty"` 
	Addresses           []interface{}  `json:"addresses,omitempty"` 
	LinkCollection
}

func (c Customer) String() string {
	return Stringify(c)
}

