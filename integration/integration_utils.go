package desk_integration

import (
	"github.com/talbright/go-desk/desk"
	"os"
	"log"
	"fmt"
)

//TODO this shouldn't be hardcoded
const DefaultCustomerId int = 192220782

func init() {
	SetupLogging()
}

func CreateClient() *desk.Client {
	siteUrl := os.Getenv("DESK_SITE_URL")
	userEmail := os.Getenv("DESK_SITE_EMAIL")
	userPassword := os.Getenv("DESK_SITE_PASS")
	return desk.NewClient(nil, siteUrl, userEmail, userPassword)
}

func SetupLogging() {
	f, err := os.OpenFile("test.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
			panic(fmt.Sprintf("error opening log file: %v", err))
	}
	log.SetOutput(f)
}

func BuildSampleMessage() *desk.Message {
	message := desk.MessageBuilder.
		SetString("Direction", "in").
		SetString("Status", "received").
		SetString("To", "someone@desk.com").
		SetString("From", "someone-else@desk.com").
		SetString("Subject", "Case created by API via desk-go").
		SetString("Body", "Please assist me with this case").
		BuildMessage()
	return &message
}

func BuildSampleDraft() *desk.Draft {
	draft := desk.NewDraft()
	draft.Body = desk.String("nice body")
	draft.Direction = desk.String("out")
	draft.Status = desk.String("draft")
	return draft
}

func BuildSampleReply() *desk.Reply {
	reply := desk.ReplyBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildReply()
	return &reply
}

func BuildSampleCase() *desk.Case {
	message := desk.MessageBuilder.
		SetString("Direction", "in").
		SetString("Status", "received").
		SetString("To", "someone@desk.com").
		SetString("From", "someone-else@desk.com").
		SetString("Subject", "Case created by API via desk-go").
		SetString("Body", "Please assist me with this case").
		BuildMessage()
	caze := desk.CaseBuilder.
		SetString("Type", "email").
		SetString("Subject", "Case created by API via desk-go").
		SetInt("Priority", 4).
		SetString("Status", "received").
		SetMessage(message).
		AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
		BuildCase()
	return &caze
}

func BuildSampleNote() *desk.Note {
	note := desk.NoteBuilder.
		SetString("Body","sexy body").
		BuildNote()
	return &note
}

