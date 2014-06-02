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
  restful := Restful{}
	msg := new(Message)
  resp,err := restful.
    Get(fmt.Sprintf("cases/%s/message", caseId)).
    Json(msg).
    Client(s.client).
    Do()
	return msg, resp, err
}

// Update the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-update
func (s *MessageService) Update(caseId string,msg *Message,params *url.Values) (*Message, *http.Response, error) {
  restful:=Restful{}
  updatedMsg:=new(Message)
  resp,err:=restful.
    Patch(fmt.Sprintf("cases/%s/message", caseId)).
    Body(updatedMsg).
    Params(params).
    Json(updatedMsg).
    Client(s.client).
    Do()
	return updatedMsg, resp, err
}

// Delete the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-delete
func (s *MessageService) Delete(caseId string) (*http.Response, error) {
  restful:=Restful{}
  resp,err:=restful.
    Delete(fmt.Sprintf("cases/%s/message",caseId)).
    Client(s.client).
    Do()
	return resp, err
}
