package desk

import (
    "fmt"
    "testing"
    "time"
    . "github.com/smartystreets/goconvey/convey"
)

func TestJsonBuilder(t *testing.T) {
    fmt.Println("")
    Convey("Builder",t,func() {
      Convey("should set string fields",func() {
        caze:=CaseBuilder.SetString("Subject","foo").BuildCase()
        So(*caze.Subject,ShouldEqual,"foo")
      })
      Convey("should set int fields",func() {
        caze:=CaseBuilder.SetInt("Id",99).BuildCase()
        So(*caze.Id,ShouldEqual,99)
      })
      Convey("should set Timestamp fields",func() {
        timet:=Timestamp{time.Now()}
        caze:=CaseBuilder.SetTimestamp("LockedUntil",timet).BuildCase()
        So(caze.LockedUntil.String(),ShouldEqual,timet.String())
      })
      Convey("should set TimestampNow fields",func() {
        //not sure how to mock this out in Go
        caze:=CaseBuilder.SetTimestampNow("LockedUntil").BuildCase()
        So(caze.LockedUntil.String(),ShouldNotBeNil)
      })
      Convey("should set Links field",func() {
        links:=make(map[string]map[string]interface{})
        caze:=CaseBuilder.SetLinks(links).BuildCase()
        So(caze.Links,ShouldNotBeNil)
      })
      Convey("should set Message field",func() {
        msg:=Message{}
        caze:=CaseBuilder.SetMessage(msg).BuildCase()
        So(caze.Message,ShouldNotBeNil)
      })
      Convey("should add CustomFields pair",func() {
        caze:=CaseBuilder.
          AddCustomField("foo1","bar1").
          AddCustomField("foo2","bar2").
          BuildCase()
        So(caze.CustomFields["foo1"],ShouldEqual,"bar1")
        So(caze.CustomFields["foo2"],ShouldEqual,"bar2")
      })
      Convey("should add address",func() {
        customer:=CustomerBuilder.AddAddress("123 somewhere","primary").BuildCustomer()
        So(customer.Addresses,ShouldNotBeNil)
        So(customer.Addresses[0]["value"],ShouldEqual,"123 somewhere")
        So(customer.Addresses[0]["type"],ShouldEqual,"primary")
      })
      Convey("should add email",func() {
        customer:=CustomerBuilder.AddEmail("me@me.com","primary").BuildCustomer()
        So(customer.Emails,ShouldNotBeNil)
        So(customer.Emails[0]["value"],ShouldEqual,"me@me.com")
        So(customer.Emails[0]["type"],ShouldEqual,"primary")
      })
      Convey("should add phone number",func() {
        customer:=CustomerBuilder.AddPhoneNumber("1231231234","primary").BuildCustomer()
        So(customer.PhoneNumbers,ShouldNotBeNil)
        So(customer.PhoneNumbers[0]["value"],ShouldEqual,"1231231234")
        So(customer.PhoneNumbers[0]["type"],ShouldEqual,"primary")
      })
      Convey("should build Customer struct",func() {
        customer:=CustomerBuilder.BuildCustomer()
        So(customer,ShouldNotBeNil)
        So(customer,ShouldHaveSameTypeAs,Customer{})
      })
      Convey("should build Case struct",func() {
        caze:=CaseBuilder.BuildCase()
        So(caze,ShouldNotBeNil)
        So(caze,ShouldHaveSameTypeAs,Case{})
      })
      Convey("should build Message struct",func() {
        msg:=MessageBuilder.BuildMessage()
        So(msg,ShouldNotBeNil)
        So(msg,ShouldHaveSameTypeAs,Message{})
      })
      Convey("should build Reply struct",func() {
        reply:=ReplyBuilder.BuildReply()
        So(reply,ShouldNotBeNil)
        So(reply,ShouldHaveSameTypeAs,Reply{})
      })
    })
}
