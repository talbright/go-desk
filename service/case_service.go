package service

import (
	"bytes"
	"encoding/json"
	. "github.com/talbright/go-desk/resource"
	"net/http"
	"net/url"
)

type CaseService struct {
	client     *Client
	Message    *MessageService
	Reply      *ReplyService
	Draft      *DraftService
	Note       *NoteService
	Attachment *AttachmentService
}

func NewCaseService(httpClient *Client) *CaseService {
	s := &CaseService{client: httpClient}
	s.Message = &MessageService{client: httpClient}
	s.Reply = &ReplyService{client: httpClient}
	s.Draft = &DraftService{client: httpClient}
	s.Note = &NoteService{client: httpClient}
	s.Attachment = &AttachmentService{client: httpClient}
	return s
}

// Get retrieves a single case by ID.
// See Desk API method show (http://dev.desk.com/API/cases/#show)
func (s *CaseService) Get(id string) (*Case, *http.Response, error) {
	restful := Restful{}
	cse := NewCase()
	path := NewIdentityResourcePath(id, cse)
	resp, err := restful.
		Get(path.Path()).
		Json(cse).
		Client(s.client).
		Do()
	return cse, resp, err
}

// List cases with filtering and pagination.
// See Desk API method list (http://dev.desk.com/API/cases/#list)
func (s *CaseService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewCase())
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

// Search for cases with filtering and pagination.
// See Desk API method list (http://dev.desk.com/API/cases/#search)
func (s *CaseService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewCase()).SetAction("search")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Query(q).
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

func (s *CaseService) Feed(id string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewCase()).SetAction("feed")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(s.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = s.unravelFeedPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (s *CaseService) History(id string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewCase()).SetAction("history")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(s.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = s.unravelHistoryPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (s *CaseService) Labels(id string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewCase()).SetAction("labels")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(s.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = s.unravelLabelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

// Create a case.(does not route through customer cases path)
// See Desk API: http://dev.desk.com/API/cases/#create
func (s *CaseService) Create(cse *Case) (*Case, *http.Response, error) {
	restful := Restful{}
	createdCase := NewCase()
	path := NewResourcePath(NewCase())
	resp, err := restful.
		Post(path.Path()).
		Body(cse).
		Json(createdCase).
		Client(s.client).
		Do()
	return createdCase, resp, err
}

// Update a case.
// See Desk API: http://dev.desk.com/API/cases/#update
func (s *CaseService) Update(cse *Case) (*Case, *http.Response, error) {
	restful := Restful{}
	updatedCase := NewCase()
	path := NewIdentityResourcePath(cse.GetResourceId(), cse)
	resp, err := restful.
		Patch(path.Path()).
		Body(cse).
		Json(updatedCase).
		Client(s.client).
		Do()
	return updatedCase, resp, err
}

// Delete a case by ID.
// See Desk API: http://dev.desk.com/API/cases/#delete
func (s *CaseService) Delete(id string) (*http.Response, error) {
	restful := Restful{}
	path := NewIdentityResourcePath(id, NewCase())
	resp, err := restful.
		Delete(path.Path()).
		Client(s.client).
		Do()
	return resp, err
}

//Forward a case
//See Desk API: http://dev.desk.com/API/cases/#forward
func (s *CaseService) Forward(id string, recipients string, note string) (*http.Response, error) {
	forward := make(map[string]string)
	forward["to"] = recipients
	forward["note_text"] = note
	restful := Restful{}
	path := NewIdentityResourcePath(id, NewCase()).SetAction("forward")
	resp, err := restful.
		Post(path.Path()).
		Client(s.client).
		Body(forward).
		Do()
	return resp, err
}

func (s *CaseService) unravelPage(page *Page) error {
	cases := new([]Case)
	err := json.Unmarshal(*page.Embedded.RawEntries, &cases)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*cases))
	for i, v := range *cases {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}

func (s *CaseService) unravelFeedPage(page *Page) error {
	var container interface{}

	decoder := json.NewDecoder(bytes.NewReader(*page.Embedded.RawEntries))
	decoder.UseNumber()
	err := decoder.Decode(&container)
	if err != nil {
		return err
	}

	feedItems := container.([]interface{})
	page.Embedded.Entries = make([]interface{}, 0)
	for _, v := range feedItems {
		entry := v.(map[string]interface{})
		links := entry["_links"].(map[string]interface{})
		self := links["self"].(map[string]interface{})
		remarshalled, err := json.Marshal(v)
		if err != nil {
			return err
		}
		switch self["class"] {
		case "note":
			note := NewNote()
			err = json.Unmarshal(remarshalled, &note)
			if err != nil {
				return err
			}
			page.Embedded.Entries = append(page.Embedded.Entries, interface{}(note))
		default:
			reply := NewReply()
			err = json.Unmarshal(remarshalled, &reply)
			if err != nil {
				return err
			}
			page.Embedded.Entries = append(page.Embedded.Entries, interface{}(reply))
		}
	}

	return err
}

func (s *CaseService) unravelHistoryPage(page *Page) error {
	caseEvents := new([]CaseEvent)
	err := json.Unmarshal(*page.Embedded.RawEntries, &caseEvents)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*caseEvents))
	for i, v := range *caseEvents {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}

func (s *CaseService) unravelLabelPage(page *Page) error {
	labels := new([]Label)
	err := json.Unmarshal(*page.Embedded.RawEntries, &labels)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*labels))
	for i, v := range *labels {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
