package integration_tests

import (
	// "log"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	resource "github.com/talbright/go-desk/resource"
	types "github.com/talbright/go-desk/types"
	"net/url"
	"testing"
	"time"
)

func TestCustomerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a customer by ID", t, func() {
		customer, _, err := client.Customer.Get(fmt.Sprintf("%d", DefaultCustomerId))
		So(err, ShouldBeNil)
		So(customer, ShouldNotBeNil)
	})

	Convey("should be able to retrieve a list of customers", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Customer.List(&listParams)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to search for customers", t, func() {
		searchParams := url.Values{}
		searchParams.Add("sort_field", "created_at")
		searchParams.Add("sort_direction", "asc")
		searchParams.Add("max_id", "200000000")
		collection, _, err := client.Customer.Search(&searchParams, nil)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to create a customer", t, func() {
		customer := resource.NewCustomer()
		customer.FirstName = types.String("James")
		customer.LastName = types.String("Dean")
		newCustomer, _, err := client.Customer.Create(customer)
		So(err, ShouldBeNil)
		So(newCustomer, ShouldNotBeNil)
	})

	Convey("should be able to update a customer", t, func() {
		background := fmt.Sprintf("background updated at %v", time.Now())
		customer := resource.NewCustomer()
		customer.Id = types.Integer(DefaultCustomerId)
		customer.Background = types.String(background)
		updatedCustomer, _, err := client.Customer.Update(customer)
		So(err, ShouldBeNil)
		So(updatedCustomer, ShouldNotBeNil)
		So(*updatedCustomer.Background, ShouldEqual, background)
	})

	Convey("should be able to retrieve cases for a customer", t, func() {
		params := url.Values{}
		params.Add("sort_field", "created_at")
		params.Add("sort_direction", "asc")
		collection, _, err := client.Customer.Cases(fmt.Sprintf("%d", DefaultCustomerId), &params)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

}
