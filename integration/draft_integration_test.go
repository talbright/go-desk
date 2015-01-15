package desk_integration

import (
	"log"
	"testing"
	"fmt"
	"time"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDraftIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to create a case draft", t, func() {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		log.Printf("created case %v",createdCase)
		draft := BuildSampleDraft()
		newDraft, _, err := client.Case.Draft.Create(createdCase.GetResourceId(),draft)
		log.Printf("draft response %v",newDraft)
		So(err,ShouldBeNil)
	})
	Convey("should be able to show a case draft",t,func() {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		log.Printf("created case %v",createdCase)
		draft := BuildSampleDraft()
		newDraft, _, err := client.Case.Draft.Create(createdCase.GetResourceId(),draft)
		log.Printf("draft response %v",newDraft)
		So(err,ShouldBeNil)
		showDraft, _, err := client.Case.Draft.Get(createdCase.GetResourceId())
		So(err,ShouldBeNil)
		So(showDraft.GetResourceId(),ShouldEqual,newDraft.GetResourceId())
		So(*showDraft.Subject,ShouldEqual,*newDraft.Subject)
	})
	Convey("should be able to update a case draft",t,func() {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		log.Printf("created case %v",createdCase)
		draft := BuildSampleDraft()
		newDraft, _, err := client.Case.Draft.Create(createdCase.GetResourceId(),draft)
		log.Printf("draft response %v",newDraft)
		updatedBody := fmt.Sprintf("body updated at %v", time.Now())
		newDraft.Body = &updatedBody
		//TODO this marshalls to null, but the API cannot handle null
		delete(newDraft.Links, "outbound_mailbox")
		updatedDraft, _, err := client.Case.Draft.Update(createdCase.GetResourceId(), newDraft)
		So(err,ShouldBeNil)
		So(*updatedDraft.Body,ShouldEqual,*newDraft.Body)
	})

}

