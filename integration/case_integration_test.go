package desk_integration

import (
	"log"
	"net/url"
	"testing"
	"fmt"
	"time"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/talbright/go-desk/desk"
)

func TestCaseIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()
	Convey("should be able to retrieve a case by ID", t, func() {
		cse, _, err := client.Case.Get("1")
		So(err,ShouldBeNil)
		log.Println("case %v",cse)
		So(*cse.Subject,ShouldNotBeBlank)
	})
	Convey("should be able to list cases",t,func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Case.List(&listParams)
		So(err,ShouldBeNil)
		log.Println("collection %v",collection)
		So(collection,ShouldHaveSameTypeAs,&desk.Page{})
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
	Convey("should be able to search for cases",t,func() {
		searchParams := url.Values{}
		searchParams.Add("sort_field", "created_at")
		searchParams.Add("sort_direction", "asc")
		searchParams.Add("status", "new")
		collection, _, err := client.Case.Search(&searchParams, nil)
		So(err,ShouldBeNil)
		log.Println("collection %v",collection)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
	Convey("should be able to update a case",t,func() {
		subject := desk.String(fmt.Sprintf("updated case at %v", time.Now()))
		cse := desk.NewCase()
		cse.Subject = subject
		cse.SetResourceId("1")
		updatedCase, _, err := client.Case.Update(cse)
		So(err,ShouldBeNil)
		log.Printf("Updated case: %v\n", updatedCase)
		So(*updatedCase.Subject,ShouldEqual,*subject)
		So(*updatedCase.Blurb,ShouldNotBeBlank)
	})
	Convey("should be able to create a case",t,func() {
		cse := BuildSampleCase()
		newCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		log.Printf("Created case: %v\n", newCase)
		So(newCase.GetResourceId(),ShouldNotBeBlank)
	})
	Convey("should be able to delete a case",t,func() {
		cse := BuildSampleCase()
		newCase, _, err := client.Case.Create(cse)
		So(err,ShouldBeNil)
		resp, err := client.Case.Delete(newCase.GetResourceId())
		log.Printf("Delete response: %v\n", resp)
		So(err,ShouldBeNil)
	})
	Convey("should be able to forward a case",t,func() {
		resp, err := client.Case.Forward("1", "someone@desk.com", "some note")
		So(err,ShouldBeNil)
		log.Printf("Forward response: %v\n", resp)
	})
	Convey("should be able to get case feed",t,func() {
		collection, _, err := client.Case.Feed("1",nil)
		So(err,ShouldBeNil)
		So(collection,ShouldNotBeNil)
		log.Println("collection %v",collection)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
	Convey("should be able to get case history",t,func() {
		collection, _, err := client.Case.History("1",nil)
		So(err,ShouldBeNil)
		So(collection,ShouldNotBeNil)
		log.Println("collection %v",collection)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
	Convey("should be able to get case labels",t,func() {
		collection, _, err := client.Case.Labels("1",nil)
		So(err,ShouldBeNil)
		So(collection,ShouldNotBeNil)
		log.Println("collection %v",collection)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
}

