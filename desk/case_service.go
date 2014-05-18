package desk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CaseService struct {
	client *Client
}

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

func (s *CaseService) List(params *url.Values) (*Collection, *http.Response, error) {
	path := fmt.Sprintf("cases")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	}
	req, err := s.client.NewRequest("GET", path, nil)
	collection := new(Collection)
	resp, err := s.client.Do(req, collection)
	if err != nil {
		return nil, resp, err
	}
	cases := new([]Case)
	err = json.Unmarshal(*collection.Embed.RawEntries, &cases)
	if err != nil {
		return nil, resp, err
	}
	collection.Embed.Entries = make([]interface{}, len(*cases))
	for i, v := range *cases {
		collection.Embed.Entries[i] = interface{}(v)
	}
	collection.Embed.RawEntries = nil
	return collection, resp, err
}

func (s *CaseService) Search(params *url.Values,q *string) (*Collection, *http.Response, error) {
	path := fmt.Sprintf("cases/search")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	} else if q != nil {
		path = fmt.Sprintf("%v?%v", path, q)
  }
	req, err := s.client.NewRequest("GET", path, nil)
	collection := new(Collection)
	resp, err := s.client.Do(req, collection)
	if err != nil {
		return nil, resp, err
	}
	cases := new([]Case)
	err = json.Unmarshal(*collection.Embed.RawEntries, &cases)
	if err != nil {
		return nil, resp, err
	}
	collection.Embed.Entries = make([]interface{}, len(*cases))
	for i, v := range *cases {
		collection.Embed.Entries[i] = interface{}(v)
	}
	collection.Embed.RawEntries = nil
	return collection, resp, err
}

func (s *CaseService) Create(cse *Case, customer *Customer, message *Message) (*Case, *http.Response, error) {
	return nil, nil, nil
}

