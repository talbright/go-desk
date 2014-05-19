package desk

type LinkCollection struct {
	Links           *map[string]interface{} `json:"_links,omitempty"`
}

func (c LinkCollection) String() string {
	return Stringify(c)
}

