package desk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ReplyService struct {
	client *Client
}

// Get retrieves a reply for a case.
// See Desk API: http://dev.desk.com/API/cases/#replies-show
func (c *ReplyService) Get(caseId string, replyId string) (*Reply, *http.Response, error) {
	restful := Restful{}
	reply := NewReply()
	replyPath := NewIdentityResourcePath(replyId,NewReply())
	casePath := NewIdentityResourcePath(caseId,NewCase()).AppendPath(replyPath)
	resp, err := restful.
		Get(casePath.Path()).
		Json(reply).
		Client(c.client).
		Do()
	return reply, resp, err
}

// List replies with filtering and pagination.
// See Desk API: http://dev.desk.com/API/cases/#replies-list
func (c *ReplyService) List(caseId string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	replyPath := NewResourcePath(NewReply())
	casePath := NewIdentityResourcePath(caseId,NewCase()).AppendPath(replyPath)
	resp, err := restful.
		Get(casePath.Path()).
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

// Create a reply.
// See Desk API: http://dev.desk.com/API/cases/#replies-create
func (c *ReplyService) Create(caseId string, reply *Reply) (*Reply, *http.Response, error) {
	restful := Restful{}
	createdReply := NewReply()
	replyPath := NewResourcePath(createdReply)
	casePath := NewIdentityResourcePath(caseId,NewCase()).AppendPath(replyPath)
	resp, err := restful.
		Post(casePath.Path()).
		Body(reply).
		Json(createdReply).
		Client(c.client).
		Do()
	return createdReply, resp, err
}

// Update a reply.
// See Desk API: http://dev.desk.com/API/replies/#update
func (c *ReplyService) Update(caseId string, reply *Reply) (*Reply, *http.Response, error) {
	restful := Restful{}
	updatedReply := NewReply()
	casePath := NewIdentityResourcePath(caseId,NewCase()).SetNested(reply)
	resp, err := restful.
		Patch(casePath.Path()).
		Body(reply).
		Json(updatedReply).
		Client(c.client).
		Do()
	return updatedReply, resp, err
}

// Delete a reply for a case.
// See Desk API: http://dev.desk.com/API/cases/#replies-show
func (c *ReplyService) Delete(caseId string, replyId string) (*http.Response, error) {
	restful := Restful{}
	replyPath := NewIdentityResourcePath(replyId,NewReply())
	casePath := NewIdentityResourcePath(caseId,NewCase()).AppendPath(replyPath)
	resp, err := restful.
		Delete(casePath.Path()).
		Client(c.client).
		Do()
	return resp, err
}

func (c *ReplyService) unravelPage(page *Page) error {
	replies := new([]Reply)
	err := json.Unmarshal(*page.Embedded.RawEntries, &replies)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*replies))
	for i, v := range *replies {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
