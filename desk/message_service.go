package desk

import (
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
	msg := NewMessage()
	path := NewIdentityResourcePath(caseId,NewCase()).SetNested(msg)
	resp, err := restful.
		Get(path.Path()).
		Json(msg).
		Client(s.client).
		Do()
	return msg, resp, err
}

// Update the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-update
func (s *MessageService) Update(caseId string, msg *Message, params *url.Values) (*Message, *http.Response, error) {
	restful := Restful{}
	updatedMsg := NewMessage()
	path := NewIdentityResourcePath(caseId,NewCase()).SetNested(NewMessage())
	resp, err := restful.
		Patch(path.Path()).
		Body(msg).
		Params(params).
		Json(updatedMsg).
		Client(s.client).
		Do()
	return updatedMsg, resp, err
}

// Delete the case message.
// See Desk API: http://dev.desk.com/API/cases/#message-delete
func (s *MessageService) Delete(caseId string) (*http.Response, error) {
	restful := Restful{}
	path := NewIdentityResourcePath(caseId,NewCase()).SetNested(NewMessage())
	resp, err := restful.
		Delete(path.Path()).
		Client(s.client).
		Do()
	return resp, err
}
