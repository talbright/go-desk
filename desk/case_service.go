package desk

import (
	"fmt"
  "net/http"
  /* "log" */
)

type CaseService struct {
	client *Client
}

func (s *CaseService) Get(id string) (*Case, *http.Response, error) {
	u := fmt.Sprintf("cases/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
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

func (s *CaseService) Create(cse *Case) (*Case, *http.Response, error) {
  return nil,nil,nil
}

