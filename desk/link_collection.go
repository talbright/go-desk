package desk

// LinkCollection contains related resources which are linked and
// embedded using the HAL specification. Most API methods include
// embedded links in their response.
// See Desk API (http://dev.desk.com/API/using-the-api/#relationships)
type LinkCollection struct {
	Links *map[string]interface{} `json:"_links,omitempty"`
}

func (c LinkCollection) String() string {
	return Stringify(c)
}
