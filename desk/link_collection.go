package desk

// LinkCollection contains related resources which are linked and
// embedded using the HAL specification. Most API methods include
// embedded links in their response.
// See Desk API (http://dev.desk.com/API/using-the-api/#relationships)
type LinkCollection struct {
	Links map[string]map[string]interface{}   `json:"_links,omitempty"`
}

func NewLinkCollection() *LinkCollection {
	c := &LinkCollection{}
  c.Links = make(map[string]map[string]interface{})
	return c
}

func (c LinkCollection) String() string {
	return Stringify(c)
}

func (c *LinkCollection) GetLinkSubItemStringValue(link string, subitem string) (string) {
  var str string
  if c.HasLinkAndSubItem(link,subitem) {
    str = c.Links[link][subitem].(string)
  }
  return str
}

func (c* LinkCollection) AddLinkSubItemStringValue(link string, subitem string, value string) {
  if c.Links == nil {
    c.Links = make(map[string]map[string]interface{})
  }  
  if c.Links[link] == nil {
    c.Links[link] = make(map[string]interface{})
  }
  c.Links[link][subitem]=value
}

func (c* LinkCollection) AddHrefLink(class string, href string) {
  c.AddLinkSubItemStringValue(class,"href",href)
  c.AddLinkSubItemStringValue(class,"class",class)
}

func (c* LinkCollection) HasLink(name string) (bool) {
  return c.Links!=nil && c.Links[name]!=nil 
}

func (c* LinkCollection) HasLinkAndSubItem(name string,subitem string) (bool) {
  return c.HasLink(name) && c.Links[name][subitem]!=nil
}

