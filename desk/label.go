package desk

type Label struct {
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Color        *string                `json:"color,omitempty"`
	Enabled      *bool                  `json:"enabled,omitempty"`
	Active       *bool                  `json:"active,omitempty"`
	Postion      *int                   `json:"position,omitempty"`
	Types        []string               `json:"types,omitempty"`
	Resource
}

func NewLabel() *Label {
	label := &Label{}
	label.InitializeResource(label)
	return label
}

func (c Label) String() string {
	return Stringify(c)
}

