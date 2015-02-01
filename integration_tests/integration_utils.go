package integration_tests

import (
	service "github.com/talbright/go-desk/service"
	resource "github.com/talbright/go-desk/resource"
	types "github.com/talbright/go-desk/types"
	"os"
	"log"
	"fmt"
)

//TODO this shouldn't be hardcoded
const DefaultCustomerId int = 192220782

func init() {
	SetupLogging()
}

func CreateClient() *service.Client {
	siteUrl := os.Getenv("resource_SITE_URL")
	userEmail := os.Getenv("resource_SITE_EMAIL")
	userPassword := os.Getenv("resource_SITE_PASS")
	return service.NewClient(nil, siteUrl, userEmail, userPassword)
}

func SetupLogging() {
	f, err := os.OpenFile("test.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
			panic(fmt.Sprintf("error opening log file: %v", err))
	}
	log.SetOutput(f)
}

func BuildSampleMessage() *resource.Message {
	message := resource.MessageBuilder.
		SetString("Direction", "in").
		SetString("Status", "received").
		SetString("To", "someone@resource.com").
		SetString("From", "someone-else@resource.com").
		SetString("Subject", "Case created by API via resource-go").
		SetString("Body", "Please assist me with this case").
		BuildMessage()
	return &message
}

func BuildSampleDraft() *resource.Draft {
	draft := resource.NewDraft()
	draft.Body = types.String("nice body")
	draft.Direction = types.String("out")
	draft.Status = types.String("draft")
	return draft
}

func BuildSampleReply() *resource.Reply {
	reply := resource.ReplyBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildReply()
	return &reply
}

func BuildSampleCase() *resource.Case {
	message := resource.MessageBuilder.
		SetString("Direction", "in").
		SetString("Status", "received").
		SetString("To", "someone@resource.com").
		SetString("From", "someone-else@resource.com").
		SetString("Subject", "Case created by API via resource-go").
		SetString("Body", "Please assist me with this case").
		BuildMessage()
	caze := resource.CaseBuilder.
		SetString("Type", "email").
		SetString("Subject", "Case created by API via resource-go").
		SetInt("Priority", 4).
		SetString("Status", "received").
		SetMessage(message).
		AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
		BuildCase()
	return &caze
}

func BuildSampleNote() *resource.Note {
	note := resource.NoteBuilder.
		SetString("Body","sexy body").
		BuildNote()
	return &note
}

func BuildSampleAttachment() *resource.Attachment {
	attach := resource.NewAttachment()
	attach.FileName = types.String("test.png")
	attach.ContentType = types.String("image/png")
	attach.SetContent("test.png")
	return attach
}

