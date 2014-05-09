package desk

import (
	"fmt"
  "net/http"
  /* "log" */
)

type CasesService struct {
	client *Client
}

type Case struct {
	ID            *int                      `json:"id,omitempty"`
	Type          *string                   `json:"type,omitempty"`
	Status        *string                   `json:"status,omitempty"`
}

func (c Case) String() string {
	return Stringify(c)
}

func (s *CasesService) Get(id string) (*Case, *http.Response, error) {
	u := fmt.Sprintf("cases/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	cse := new(Case)
	resp, err := s.client.Do(req, cse)
  /* var v map[string]interface{} */
	/* resp, err := s.client.Do(req, &v) */
  /* log.Printf("response %v",v)  */
	if err != nil {
		return nil, resp, err
	}
	return cse, resp, err
}

