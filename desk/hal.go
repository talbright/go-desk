package desk

import (
	"fmt"
	"strconv"
	"strings"
)

type Hal struct {
	Id    *int                              `json:"id,omitempty"`
	Links map[string]map[string]interface{} `json:"_links,omitempty"`
}

func NewHal() *Hal {
	c := &Hal{}
	c.Links = make(map[string]map[string]interface{})
	return c
}

func (c Hal) String() string {
	return Stringify(c)
}

func (c *Hal) GetId() int {
	if c.Id == nil || *c.Id == 0 {
		idLink := c.GetHrefLink("self")
		if idLink != "" {
			sections := strings.Split(idLink, "/")
			if len(sections) > 0 {
				strId := sections[len(sections)-1]
				intId, _ := strconv.Atoi(strId)
				c.SetId(intId)
			}
		} else {
			c.SetId(-1)
		}
	}
	return *c.Id
}

func (c *Hal) SetId(id int) {
	c.Id = new(int)
	*c.Id = id
}

func (c *Hal) GetLinkSubItemStringValue(link string, subitem string) string {
	var str string
	if c.HasLinkAndSubItem(link, subitem) {
		str = c.Links[link][subitem].(string)
	}
	return str
}

func (c *Hal) AddLinkSubItemStringValue(link string, subitem string, value string) {
	if c.Links == nil {
		c.Links = make(map[string]map[string]interface{})
	}
	if c.Links[link] == nil {
		c.Links[link] = make(map[string]interface{})
	}
	c.Links[link][subitem] = value
}

func (c *Hal) AddHrefLink(class string, href string) {
	c.AddLinkSubItemStringValue(class, "href", href)
	c.AddLinkSubItemStringValue(class, "class", class)
}

func (c *Hal) GetHrefLink(class string) string {
	var href string
	if c.HasLinkAndSubItem(class, "href") {
		href = c.Links["self"]["href"].(string)
	}
	return href
}

func (c *Hal) HasLink(name string) bool {
	return c.Links != nil && c.Links[name] != nil
}

func (c *Hal) HasLinkAndSubItem(name string, subitem string) bool {
	return c.HasLink(name) && c.Links[name][subitem] != nil
}

func (c *Hal) GetResourceId() string {
	var id string
	if c.GetId()>=0 {
		id = fmt.Sprintf("%d", c.GetId())
	}
	return id
}

func (c *Hal) SetResourceId(id string) {
	intId, _ := strconv.Atoi(id)
	c.SetId(intId)
}

func (c *Hal) GetResourcePath(resource Resourceful, options ...func(*ResourcePath)) ResourcePath {
	pathing:=ResourcePath{Target: resource}
	pathing.SetMember()
	for _,option := range options {
		option(&pathing)
	}
	return pathing
}

