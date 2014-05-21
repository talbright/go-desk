package desk

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "net/url"
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
func (s *MessageService) Update(caseId string, msg *Message) (*Message, *http.Response, error) {
	return nil, nil, nil
}

// Delete the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-delete
func (s *MessageService) Delete(caseId string) (*http.Response, error) {
	return nil, nil
}
