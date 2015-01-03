package desk

import (
	"encoding/json"
	"fmt"
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
	reply := new(Reply)
	resp, err := restful.
		Get(fmt.Sprintf("cases/%s/replies/%s", caseId, replyId)).
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
	resp, err := restful.
		Get(fmt.Sprintf("cases/%s/replies", caseId)).
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
	createdReply := new(Reply)
	resp, err := restful.
		Post(fmt.Sprintf("cases/%s/replies", caseId)).
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
	updatedReply := new(Reply)
	resp, err := restful.
		Patch(fmt.Sprintf("cases/%s/replies/%d", caseId, reply.GetId())).
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
	resp, err := restful.
		Delete(fmt.Sprintf("cases/%s/replies/%s", caseId, replyId)).
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
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
