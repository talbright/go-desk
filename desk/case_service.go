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
}

func NewCaseService(httpClient *Client) *CaseService {
	s := &CaseService{client: httpClient}
	s.Message = &MessageService{client: httpClient}
	return s
}

// Get retrieves a single case by ID.
// See Desk API method show (http://dev.desk.com/API/cases/#show)
func (s *CaseService) Get(id string) (*Case, *http.Response, error) {
	path := fmt.Sprintf("cases/%v", id)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	cse := new(Case)
	resp, err := s.client.Do(req, cse)
	//For posterity sake leaving here...an example of making a request
	//and dumping it into a generic map
	//var v map[string]interface{}
	//resp, err := s.client.Do(req, &v)
	//log.Printf("response %v",v)
	if err != nil {
		return nil, resp, err
	}
	return cse, resp, err
}

// List cases with filtering and pagination.
// See Desk API method list (http://dev.desk.com/API/cases/#list)
func (s *CaseService) List(params *url.Values) (*Page, *http.Response, error) {
	path := fmt.Sprintf("cases")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	}
	req, err := s.client.NewRequest("GET", path, nil)
	page := new(Page)
	resp, err := s.client.Do(req, page)
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
	path := fmt.Sprintf("cases/search")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	} else if q != nil {
		path = fmt.Sprintf("%v?%v", path, q)
	}
	req, err := s.client.NewRequest("GET", path, nil)
	page := new(Page)
	resp, err := s.client.Do(req, page)
	if err != nil {
		return nil, resp, err
	}
  err = s.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
	return page, resp, err
}

// Create a case.
// See Desk API: http://dev.desk.com/API/cases/#create
func (s *CaseService) Create(cse *Case, customer *Customer, message *Message) (*Case, *http.Response, error) {
	u := fmt.Sprintf("cases")
	req, err := s.client.NewRequest("POST", u, cse)
	if err != nil {
		return nil, nil, err
	}

	c := new(Case)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// Update a case.
// See Desk API: http://dev.desk.com/API/cases/#update
func (s *CaseService) Update(cse *Case) (*Case, *http.Response, error) {
	u := fmt.Sprintf("cases/%d", *cse.ID)
	req, err := s.client.NewRequest("PATCH", u, cse)
	if err != nil {
		return nil, nil, err
	}

	c := new(Case)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// Delete a case by ID.
// See Desk API: http://dev.desk.com/API/cases/#delete
func (s *CaseService) Delete(id string) (*http.Response, error) {
	return nil, nil
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

