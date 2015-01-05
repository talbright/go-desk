package desk

import (
	"net/http"
)

type DraftService struct {
	client *Client
}

// Get retrieves the draft reply for a case.
// See Desk API: http://dev.desk.com/API/cases/#drafts-show
func (c *DraftService) Get(id string) (*Draft, *http.Response, error) {
	restful := Restful{}
	draft := NewDraft()
	path := NewIdentityResourcePath(id,NewCase()).SetAction("replies").SetNested(NewDraft())
	resp, err := restful.
		Get(path.Path()).
		Json(draft).
		Client(c.client).
		Do()
	return draft, resp, err
}

// Create a draft.
// See Desk API: http://dev.desk.com/API/cases/#drafts-create
func (c *DraftService) Create(id string, draft *Draft) (*Draft, *http.Response, error) {
	restful := Restful{}
	createdDraft := NewDraft()
	path := NewIdentityResourcePath(id,NewCase()).SetAction("replies").SetNested(draft)
	resp, err := restful.
		Post(path.Path()).
		Body(draft).
		Json(createdDraft).
		Client(c.client).
		Do()
	return createdDraft, resp, err
}

// Update a draft.
// See Desk API: http://dev.desk.com/API/replies/#update
func (c *DraftService) Update(id string, draft *Draft) (*Draft, *http.Response, error) {
	restful := Restful{}
	updatedDraft := NewDraft()
	repliesPath := NewIdentityResourcePath(draft.GetResourceId(),NewReply())
	casesPath := NewIdentityResourcePath(id,NewCase()).SetAction("replies").SetSuffix(repliesPath.Path())
	resp, err := restful.
		Patch(casesPath.Path()).
		Body(draft).
		Json(updatedDraft).
		Client(c.client).
		Do()
	return updatedDraft, resp, err
}
