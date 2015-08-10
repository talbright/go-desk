package service

import (
	"encoding/json"
	. "github.com/talbright/go-desk/resource"
	"net/http"
	"net/url"
)

type UserService struct {
	client *Client
}

// Get retrieves a user.
// See Desk API: http://dev.desk.com/API/users/#show
func (c *UserService) Get(id string) (*User, *http.Response, error) {
	restful := Restful{}
	user := NewUser()
	path := NewIdentityResourcePath(id, user)
	resp, err := restful.
		Get(path.Path()).
		Json(user).
		Client(c.client).
		Do()
	return user, resp, err
}

// List users with filtering and pagination.
// See Desk API: http://dev.desk.com/API/users/#list
func (c *UserService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewUser())
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

// Update a user.
// See Desk API: http://dev.desk.com/API/users/#update
func (c *UserService) Update(user *User) (*User, *http.Response, error) {
	restful := Restful{}
	updatedUser := new(User)
	path := NewIdentityResourcePath(user.GetResourceId(), user)
	resp, err := restful.
		Patch(path.Path()).
		Body(user).
		Json(updatedUser).
		Client(c.client).
		Do()
	return updatedUser, resp, err
}

func (c *UserService) unravelPage(page *Page) error {
	users := new([]User)
	err := json.Unmarshal(*page.Embedded.RawEntries, &users)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*users))
	for i, v := range *users {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
