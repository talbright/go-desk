package integration_tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	resource "github.com/talbright/go-desk/resource"
	"log"
	"testing"
	"time"
)

func TestMessageIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to show a case message", t, func() {
		msg, _, err := client.Case.Message.Get("1")
		So(err, ShouldBeNil)
		log.Printf("retrieved message %v", msg)
		So(*msg.Body, ShouldNotBeBlank)
	})

	Convey("should be able to update a case message", t, func() {
		message := resource.MessageBuilder.
			SetString("Direction", "out").
			SetString("Status", "draft").
			SetString("To", "someone@desk.com").
			SetString("From", "someone-else@desk.com").
			SetString("Subject", "Case created by API via desk-go").
			SetString("Body", "Request for assistance denied").
			BuildMessage()
		caze := resource.CaseBuilder.
			SetString("Type", "email").
			SetString("Subject", "Case created by API via desk-go").
			SetInt("Priority", 4).
			SetString("Status", "received").
			SetMessage(message).
			AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
			BuildCase()
		newCase, _, err := client.Case.Create(&caze)
		So(err, ShouldBeNil)
		subject := fmt.Sprintf("Case updated by API via desk-go at %v", time.Now())
		updateMsg := resource.MessageBuilder.
			SetString("Subject", subject).
			BuildMessage()
		newMsg, _, err := client.Case.Message.Update(newCase.GetResourceId(), &updateMsg, nil)
		So(err, ShouldBeNil)
		So(*newMsg.Subject, ShouldEqual, subject)
	})

	Convey("should be able to delete a case message", t, func() {
		message := resource.MessageBuilder.
			SetString("Direction", "out").
			SetString("Status", "draft").
			SetString("To", "someone@desk.com").
			SetString("From", "someone-else@desk.com").
			SetString("Subject", "Case created by API via desk-go").
			SetString("Body", "Request for assistance denied").
			BuildMessage()
		caze := resource.CaseBuilder.
			SetString("Type", "email").
			SetString("Subject", "Case created by API via desk-go").
			SetInt("Priority", 4).
			SetString("Status", "received").
			SetMessage(message).
			AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
			BuildCase()
		newCase, _, err := client.Case.Create(&caze)
		So(err, ShouldBeNil)
		_, err1 := client.Case.Message.Delete(newCase.GetResourceId())
		So(err1, ShouldBeNil)
	})

}
