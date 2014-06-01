package desk

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type MessageService struct {
	client *Client
}

// Get retrieves the message for a case.
// See Desk API: http://dev.desk.com/API/cases/#message-show
func (s *MessageService) Get(caseId string) (*Message, *http.Response, error) {
	path := fmt.Sprintf("cases/%v/message", caseId)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	cse := new(Message)
	resp, err := s.client.Do(req, cse)
	if err != nil {
		return nil, resp, err
	}
	return cse, resp, err
}

// Update the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-update
func (s *MessageService) Update(caseId string, msg *Message,params *url.Values) (*Message, *http.Response, error) {
  path:=fmt.Sprintf("cases/%s/message",caseId)
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	}
  req,err:=s.client.NewRequest("PATCH",path,msg)
  if err !=nil {
    return nil,nil,err
  }
  m:=new(Message)
  resp,err:=s.client.Do(req,m)
	return m, resp, err
}

// Delete the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-delete
func (s *MessageService) Delete(caseId string) (*http.Response, error) {
	return nil, nil
}
