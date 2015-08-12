package integration_tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	types "github.com/talbright/go-desk/types"
	"net/url"
	"testing"
	"time"
)

func TestReplyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a list of case replies", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Case.Reply.List("1", &listParams)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able create a case reply", t, func() {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err, ShouldBeNil)
		reply := BuildSampleReply()
		createdReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), reply)
		So(err, ShouldBeNil)
		So(createdReply, ShouldNotBeNil)
		So(*createdReply.Body, ShouldEqual, "some body")
	})

	Convey("should be able to update a case reply", t, func() {
		body := types.String(fmt.Sprintf("updated body at %v", time.Now()))
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err, ShouldBeNil)
		reply := BuildSampleReply()
		createdReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), reply)
		So(err, ShouldBeNil)
		createdReply.Body = body
		updatedReply, _, err := client.Case.Reply.Update(createdCase.GetResourceId(), createdReply)
		So(err, ShouldBeNil)
		So(updatedReply, ShouldNotBeNil)
		So(*createdReply.Body, ShouldEqual, *body)
	})

	Convey("should be able to delete a case reply", t, func() {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err, ShouldBeNil)
		reply := BuildSampleReply()
		createdReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), reply)
		So(err, ShouldBeNil)
		_, err = client.Case.Reply.Delete(createdCase.GetResourceId(), createdReply.GetResourceId())
		So(err, ShouldBeNil)
	})

}
