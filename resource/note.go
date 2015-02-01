package resource

import(
	. "github.com/talbright/go-desk/types"
)

type Note struct {
	Body             *string    `json:"body,omitempty"`
	SuppressRules    *bool      `json:"supress_rules,omitempty"`
	ErasedAt         *Timestamp `json:"erased_at,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp `json:"updated_at,omitempty"`
	Resource
}

func NewNote() *Note {
	note := &Note{}
	note.InitializeResource(note)
	return note
}

func (c Note) String() string {
	return Stringify(c)
}
