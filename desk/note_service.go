package desk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type NoteService struct {
	client  *Client
}

func NewNoteService(httpClient *Client) *NoteService {
	s := &NoteService{client: httpClient}
	return s
}

func (s *NoteService) Get(caseId string,noteId string) (*Note, *http.Response, error) {
	restful := Restful{}
	note := NewNote()
	notePath := NewIdentityResourcePath(noteId,NewNote())
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(notePath)
	resp, err := restful.
		Get(path.Path()).
		Json(note).
		Client(s.client).
		Do()
	return note, resp, err
}

func (s *NoteService) Create(caseId string,note *Note) (*Note, *http.Response, error) {
	restful := Restful{}
	createdNote := NewNote()
	path := NewIdentityResourcePath(caseId,NewCase()).SetNested(createdNote)
	resp, err := restful.
		Post(path.Path()).
		Body(note).
		Json(createdNote).
		Client(s.client).
		Do()
	return createdNote, resp, err
}

func (s *NoteService) Update(caseId string,note *Note) (*Note, *http.Response, error) {
	restful := Restful{}
	updatedNote := NewNote()
	notePath := NewResourcePath(note).SetMember()
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(notePath)
	resp, err := restful.
		Patch(path.Path()).
		Body(note).
		Json(updatedNote).
		Client(s.client).
		Do()
	return updatedNote, resp, err
}

func (s *NoteService) List(caseId string,params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(caseId,NewCase()).SetNested(NewNote())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(s.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = s.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (s *NoteService) Delete(caseId string, noteId string) (*http.Response, error) {
	restful := Restful{}
	notePath := NewIdentityResourcePath(noteId,NewNote()).SetMember()
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(notePath)
	resp, err := restful.
		Delete(path.Path()).
		Client(s.client).
		Do()
	return resp, err
}

func (s *NoteService) unravelPage(page *Page) error {
	notes := new([]Case)
	err := json.Unmarshal(*page.Embedded.RawEntries, &notes)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*notes))
	for i, v := range *notes {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}

