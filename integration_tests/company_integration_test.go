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

func TestCompanyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a company by ID", t, func() {
		company, _, err := client.Company.Get(fmt.Sprintf("%d", DefaultCompanyId))
		So(err, ShouldBeNil)
		So(company, ShouldNotBeNil)
	})

	Convey("should be able to retrieve a list of companies", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Company.List(&listParams)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to search for companies", t, func() {
		searchParams := url.Values{}
		searchParams.Add("q", "desk.com")
		searchParams.Add("sort_field", "created_at")
		searchParams.Add("sort_direction", "asc")
		collection, _, err := client.Company.Search(&searchParams, nil)
		So(err, ShouldBeNil)
		log.Println("collection %v", collection)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to update a company", t, func() {
		subject := types.String(fmt.Sprintf("updated company at %v", time.Now()))
		company := resource.NewCompany()
		company.Name = subject
		company.SetResourceId(fmt.Sprintf("%d", DefaultCompanyId))
		updatedCompany, _, err := client.Company.Update(company)
		So(err, ShouldBeNil)
		log.Printf("Updated company: %v\n", updatedCompany)
		So(*updatedCompany.Name, ShouldEqual, *subject)
		So(updatedCompany.Domains[0], ShouldNotBeBlank)
	})

	Convey("should be able to create a company", t, func() {
		company := BuildSampleCompany()
		newCompany, _, err := client.Company.Create(company)
		So(err, ShouldBeNil)
		log.Printf("Created company: %v\n", newCompany)
		So(newCompany.GetResourceId(), ShouldNotBeBlank)
	})

	Convey("should be able to get company cases", t, func() {
		params := url.Values{}
		params.Add("sort_field", "created_at")
		params.Add("sort_direction", "asc")
		collection, _, err := client.Company.Cases(fmt.Sprintf("%d", DefaultCompanyId), &params)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to get company customers", t, func() {
		params := url.Values{}
		params.Add("sort_field", "created_at")
		params.Add("sort_direction", "asc")
		collection, _, err := client.Company.Customers(fmt.Sprintf("%d", DefaultCompanyId), &params)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})
}
