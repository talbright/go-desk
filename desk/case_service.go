package desk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CaseService struct {
	client      *Client
  Message     *MessageService 
  Reply       *ReplyService
  Draft       *DraftService
}

func NewCaseService(httpClient *Client) *CaseService {
	s := &CaseService{client: httpClient}
	s.Message = &MessageService{client: httpClient}
	s.Reply = &ReplyService{client: httpClient}
  s.Draft = &DraftService{client: httpClient}
	return s
}

// Get retrieves a single case by ID.
// See Desk API method show (http://dev.desk.com/API/cases/#show)
func (s *CaseService) Get(id string) (*Case, *http.Response, error) {
  restful := Restful{}
	cse := new(Case)
  resp,err := restful.
    Get(fmt.Sprintf("cases/%v", id)).
    Json(cse).
    Client(s.client).
    Do()
	return cse, resp, err
}

// List cases with filtering and pagination.
// See Desk API method list (http://dev.desk.com/API/cases/#list)
func (s *CaseService) List(params *url.Values) (*Page, *http.Response, error) {
  restful := Restful{}
  page := new(Page)
  resp,err := restful.
    Get("cases").
    Json(page).
    Params(params).
    Client(s.client).
    Do()
	if err != nil {
		return nil,resp,err
	}
  err = s.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
	return page, resp, err
}

// Search for cases with filtering and pagination.
// See Desk API method list (http://dev.desk.com/API/cases/#search)
func (s *CaseService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
  restful := Restful{}
  page := new(Page)
  resp,err := restful.
    Get("cases/search").
    Json(page).
    Query(q).
    Params(params).
    Client(s.client).
    Do()
	if err != nil {
		return nil,resp,err
	}
  err = s.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
	return page, resp, err
}

// Create a case.(does not route through customer cases path)
// See Desk API: http://dev.desk.com/API/cases/#create
func (s *CaseService) Create(cse *Case) (*Case, *http.Response, error) {
  restful:=Restful{}
  createdCase:=new(Case)
  resp,err:=restful.
    Post("cases").
    Body(cse).
    Json(createdCase).
    Client(s.client).
    Do()
  return createdCase,resp,err
}

// Update a case.
// See Desk API: http://dev.desk.com/API/cases/#update
func (s *CaseService) Update(cse *Case) (*Case, *http.Response, error) {
  restful:=Restful{}
  updatedCase:=new(Case)
  resp,err:=restful.
    Patch(fmt.Sprintf("cases/%d", cse.GetId())).
    Body(cse).
    Json(updatedCase).
    Client(s.client).
    Do()
	return updatedCase, resp, err
}

// Delete a case by ID.
// See Desk API: http://dev.desk.com/API/cases/#delete
func (s *CaseService) Delete(id string) (*http.Response, error) {
  restful:=Restful{}
  resp,err:=restful.
    Delete(fmt.Sprintf("cases/%s",id)).
    Client(s.client).
    Do()
	return resp, err
}

//Forward a case
//See Desk API: http://dev.desk.com/API/cases/#forward
func (s* CaseService) Forward(id string,recipients string,note string) (*http.Response, error) {
  forward:=make(map[string]string)
  forward["to"] = recipients
  forward["note_text"] = note
  restful:=Restful{}
  resp,err:=restful.
    Post(fmt.Sprintf("cases/%s/forward",id)).
    Client(s.client).
    Body(forward).
    Do()
	return resp, err
}

func (s* CaseService)unravelPage(page *Page) (error) {
  cases := new([]Case)
  err := json.Unmarshal(*page.Embedded.RawEntries,&cases)
  if err!= nil {
    return err
  }
  page.Embedded.Entries = make([]interface{},len(*cases))
  for i, v:= range *cases {
    page.Embedded.Entries[i] = interface{}(v)
  }
  page.Embedded.RawEntries=nil
  return err
}

