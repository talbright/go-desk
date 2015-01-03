package desk

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCustomer(t *testing.T) {
	fmt.Println("")
	Convey("AddAddress", t, func() {
		Convey("when no address exists", func() {
			Convey("should add new entry", func() {
				customer := Customer{}
				customer.AddAddress("val1", "val1Type")
				So(customer.Addresses[0], ShouldNotBeNil)
				So(customer.Addresses[0]["type"], ShouldEqual, "val1Type")
				So(customer.Addresses[0]["value"], ShouldEqual, "val1")
			})
		})
		Convey("when another address exists", func() {
			Convey("should append new entry", func() {
				customer := Customer{}
				customer.AddEmail("me@me.com", "work")
				customer.AddEmail("me2@me.com", "home")
				So(customer.Emails[1], ShouldNotBeNil)
				So(customer.Emails[1]["type"], ShouldEqual, "home")
				So(customer.Emails[1]["value"], ShouldEqual, "me2@me.com")
			})
		})
	})
	Convey("AddEmail", t, func() {
		Convey("when no email exists", func() {
			Convey("should add new entry", func() {
				customer := Customer{}
				customer.AddEmail("me@me.com", "work")
				So(customer.Emails[0], ShouldNotBeNil)
				So(customer.Emails[0]["type"], ShouldEqual, "work")
				So(customer.Emails[0]["value"], ShouldEqual, "me@me.com")
			})
		})
		Convey("when another email exists", func() {
			Convey("should append new entry", func() {
				customer := Customer{}
				customer.AddEmail("me@me.com", "work")
				customer.AddEmail("me2@me.com", "home")
				So(customer.Emails[1], ShouldNotBeNil)
				So(customer.Emails[1]["type"], ShouldEqual, "home")
				So(customer.Emails[1]["value"], ShouldEqual, "me2@me.com")
			})
		})
	})
	Convey("AddPhoneNumber", t, func() {
		Convey("when no phone number exists", func() {
			Convey("should add new entry", func() {
				customer := Customer{}
				customer.AddPhoneNumber("val1", "type1")
				So(customer.PhoneNumbers[0], ShouldNotBeNil)
				So(customer.PhoneNumbers[0]["type"], ShouldEqual, "type1")
				So(customer.PhoneNumbers[0]["value"], ShouldEqual, "val1")
			})
		})
		Convey("when another phone number exists", func() {
			Convey("should append new entry", func() {
				customer := Customer{}
				customer.AddPhoneNumber("val1", "type1")
				customer.AddPhoneNumber("val2", "type2")
				So(customer.PhoneNumbers[1], ShouldNotBeNil)
				So(customer.PhoneNumbers[1]["type"], ShouldEqual, "type2")
				So(customer.PhoneNumbers[1]["value"], ShouldEqual, "val2")
			})
		})
	})
}
