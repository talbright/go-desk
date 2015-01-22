package desk

import (
	"net/http"
	"encoding/json"
)

type AttachmentService struct {
	client *Client
}

func (s *AttachmentService) Get(caseId string,attachId string) (*Attachment, *http.Response, error) {
	restful := Restful{}
	attach := NewAttachment()
	attachPath := NewIdentityResourcePath(attachId,NewAttachment())
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(attachPath)
	resp, err := restful.
		Get(path.Path()).
		Json(attach).
		Client(s.client).
		Do()
	return attach, resp, err
}

func (s *AttachmentService) Create(caseId string, attach *Attachment) (*Attachment, *http.Response, error) {
	restful := Restful{}
	createdAttachment := NewAttachment()
	attachmentPath := NewResourcePath(createdAttachment)
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(attachmentPath)
	resp, err := restful.
		Post(path.Path()).
		Body(attach).
		Json(createdAttachment).
		Client(s.client).
		Do()
	return createdAttachment, resp, err
}

func (s *AttachmentService) Delete(caseId string,attachId string) (*http.Response, error) {
	restful := Restful{}
	attachPath := NewIdentityResourcePath(attachId,NewAttachment())
	path := NewIdentityResourcePath(caseId,NewCase()).AppendPath(attachPath)
	resp, err := restful.
		Delete(path.Path()).
		Client(s.client).
		Do()
	return resp, err
}

func (s *AttachmentService) List(caseId string) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(caseId,NewCase()).SetAction("attachments")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
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

func (s *AttachmentService) unravelPage(page *Page) error {
	attachments := new([]Attachment)
	err := json.Unmarshal(*page.Embedded.RawEntries, &attachments)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*attachments))
	for i, v := range *attachments {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}

