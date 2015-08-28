package service

import (
	"encoding/json"
	. "github.com/talbright/go-desk/resource"
	"net/http"
	"net/url"
)

type GroupService struct {
	client *Client
}

func NewGroupService(httpClient *Client) *GroupService {
	s := &GroupService{client: httpClient}
	return s
}

// Get retrieves a group.
// See Desk API: http://dev.desk.com/API/groups/#show
func (c *GroupService) Get(id string) (*Group, *http.Response, error) {
	restful := Restful{}
	group := NewGroup()
	path := NewIdentityResourcePath(id, group)
	resp, err := restful.
		Get(path.Path()).
		Json(group).
		Client(c.client).
		Do()
	return group, resp, err
}

// List group with filtering and pagination.
// See Desk API: http://dev.desk.com/API/groups/#list
func (c *GroupService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewGroup())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (c *GroupService) Users(id string) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewGroup()).SetAction("users")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelUserPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (c *GroupService) unravelPage(page *Page) error {
	group := new([]Group)
	err := json.Unmarshal(*page.Embedded.RawEntries, &group)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*group))
	for i, v := range *group {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}

func (c *GroupService) unravelUserPage(page *Page) error {
	user := new([]User)
	err := json.Unmarshal(*page.Embedded.RawEntries, &user)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*user))
	for i, v := range *user {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
