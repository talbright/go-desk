package desk

import (
	"fmt"
  "net/http"
)

type CasesService struct {
	client *Client
}

type Case struct {
	ID            *string                   `json:"id,omitempty"`
	Type          *string                   `json:"type,omitempty"`
	Status        *string                   `json:"status,omitempty"`
}

func (s *CasesService) Get(id string) (*Case, *http.Response, error) {
	u := fmt.Sprintf("cases/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	cse := new(Case)
	resp, err := s.client.Do(req, cse)
	if err != nil {
		return nil, resp, err
	}
	return cse, resp, err
}

