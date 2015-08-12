package integration_tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	resource "github.com/talbright/go-desk/resource"
	types "github.com/talbright/go-desk/types"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestNoteIntegration(t *testing.T) {

	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	createCaseWithNote := func() (*resource.Case, *resource.Note) {
		cse := BuildSampleCase()
		createdCase, _, err := client.Case.Create(cse)
		So(err, ShouldBeNil)
		note := BuildSampleNote()
		createdNote, _, err := client.Case.Note.Create(createdCase.GetResourceId(), note)
		log.Printf("note response %v", createdNote)
		So(err, ShouldBeNil)
		return createdCase, createdNote
	}

	Convey("should be able to create a case note", t, func() {
		createCaseWithNote()
	})

	Convey("should be able to show a case note", t, func() {
		cse, note := createCaseWithNote()
		showNote, _, err := client.Case.Note.Get(cse.GetResourceId(), note.GetResourceId())
		So(err, ShouldBeNil)
		So(showNote, ShouldNotBeNil)
		So(showNote.GetResourceId(), ShouldEqual, note.GetResourceId())
		So(*showNote.Body, ShouldEqual, *note.Body)
	})

	Convey("should be able to update a case note", t, func() {
		cse, note := createCaseWithNote()
		body := fmt.Sprintf("body updated at %v", time.Now())
		note.Body = types.String(body)
		updatedNote, _, err := client.Case.Note.Update(cse.GetResourceId(), note)
		So(err, ShouldBeNil)
		So(updatedNote, ShouldNotBeNil)
		So(*updatedNote.Body, ShouldEqual, body)
	})

	Convey("should be able to delete a case note", t, func() {
		cse, note := createCaseWithNote()
		_, err := client.Case.Note.Delete(cse.GetResourceId(), note.GetResourceId())
		So(err, ShouldBeNil)
	})

	Convey("should be able to list case notes", t, func() {
		cse, _ := createCaseWithNote()
		params := url.Values{}
		params.Add("sort_field", "created_at")
		params.Add("sort_direction", "asc")
		collection, _, err := client.Case.Note.List(cse.GetResourceId(), &params)
		So(err, ShouldBeNil)
		So(collection, ShouldNotBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

}
