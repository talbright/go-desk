package desk

import (
  "net/http"
  "fmt"
)

type DraftService struct {
	client *Client
}

// Get retrieves the draft reply for a case.
// See Desk API: http://dev.desk.com/API/cases/#drafts-show 
func (c* DraftService) Get(caseId string) (*Draft, *http.Response, error) {
  restful := Restful{}
  draft := new(Draft)
  resp, err := restful.
    Get(fmt.Sprintf("cases/%s/replies/draft",caseId)).
    Json(draft).
    Client(c.client).
    Do()
  return draft, resp, err
}

// Create a draft.
// See Desk API: http://dev.desk.com/API/cases/#drafts-create
func (c *DraftService) Create(caseId string,draft *Draft) (*Draft, *http.Response, error) {
  restful:=Restful{}
	createdDraft := new(Draft)
  resp,err:=restful.
    Post(fmt.Sprintf("cases/%s/replies/draft",caseId)).
    Body(draft).
    Json(createdDraft).
    Client(c.client).
    Do()
  return createdDraft,resp,err
}

// Update a draft.
// See Desk API: http://dev.desk.com/API/replies/#update
func (c *DraftService) Update(caseId string,draft *Draft) (*Draft, *http.Response, error) {
  restful:=Restful{}
  updatedDraft:=new(Draft)
  resp,err:=restful.
    Patch(fmt.Sprintf("cases/%s/replies/%d",caseId,draft.GetId())).
    Body(draft).
    Json(updatedDraft).
    Client(c.client).
    Do()
  return updatedDraft,resp,err
}

