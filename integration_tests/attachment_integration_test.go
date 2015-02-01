package integration_tests

import (
	"log"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	resource "github.com/talbright/go-desk/resource"
)

func TestAttachmentIntegration(t *testing.T) {

	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	createCaseWithAttachment:= func() (*resource.Case,*resource.Attachment) {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		attach := BuildSampleAttachment()
		createdAttachment, _, err := client.Case.Attachment.Create(createdCase.GetResourceId(),attach)
		log.Printf("attachment response %v",createdAttachment)
		So(err,ShouldBeNil)
		return createdCase,createdAttachment
	}

	Convey("should be able to create a case attachment", t, func() {
		createCaseWithAttachment()
	})

	Convey("should be able to show an attachment",t,func() {
		cse,attach := createCaseWithAttachment()
		showAttach, _, err := client.Case.Attachment.Get(cse.GetResourceId(),attach.GetResourceId())
		So(err,ShouldBeNil)
		So(showAttach,ShouldNotBeNil)
		So(showAttach.GetResourceId(),ShouldEqual,attach.GetResourceId())
		So(*showAttach.URL,ShouldNotBeNil)
	})

	Convey("should be able to delete an attachment",t,func() {
		cse,attach := createCaseWithAttachment()
		_, err := client.Case.Attachment.Delete(cse.GetResourceId(),attach.GetResourceId())
		So(err,ShouldBeNil)
	})

	Convey("should be able to list case attachments",t,func() {
		cse,_ := createCaseWithAttachment()
		collection, _, err := client.Case.Attachment.List(cse.GetResourceId())
		So(err,ShouldBeNil)
		So(collection,ShouldNotBeNil)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})

}

