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
        caze:=CaseBuilder.SetInt("ID",99).BuildCase()
        So(*caze.ID,ShouldEqual,99)
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
      Convey("should set LinkCollection field",func() {
        coll:=LinkCollection{}
        caze:=CaseBuilder.SetLinkCollection(coll).BuildCase()
        So(caze.LinkCollection,ShouldNotBeNil)
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
      Convey("should add Href link in LinkCollection",func() {
        caze:=CaseBuilder.
          AddHrefLink("customer1","/api/v2/customer/1234").
          AddHrefLink("customer2","/api/v2/customer/1234").
          BuildCase()
        So(caze.LinkCollection.Links["customer1"],ShouldNotBeNil)
        So(caze.LinkCollection.Links["customer1"]["href"],ShouldEqual,"/api/v2/customer/1234")
        So(caze.LinkCollection.Links["customer1"]["class"],ShouldEqual,"customer1")
        So(caze.LinkCollection.Links["customer2"],ShouldNotBeNil)
        So(caze.LinkCollection.Links["customer2"]["href"],ShouldEqual,"/api/v2/customer/1234")
        So(caze.LinkCollection.Links["customer2"]["class"],ShouldEqual,"customer2")
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
      Convey("should build LinkCollection struct",func() {
        links:=LinkCollectionBuilder.BuildLinkCollection()
        So(links,ShouldNotBeNil)
        So(links,ShouldHaveSameTypeAs,LinkCollection{})
      })
    })
}
