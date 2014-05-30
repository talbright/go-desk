package desk

import (
  "fmt"
  "testing"
  "time"
  . "github.com/smartystreets/goconvey/convey"
)

func TestCase(t *testing.T) {
  fmt.Println("")
  Convey("Builder", t, func() {
    Convey("builds string fields",func() {
      caze:=CaseBuilder.SetString("Type","this is my type").Build()
      So(*caze.Type,ShouldEqual,"this is my type")
    })
    Convey("builds int fields",func() {
      caze:=CaseBuilder.SetInt("ID",7).Build()
      So(*caze.ID,ShouldEqual,7)
    })
    Convey("builds Timestamp fields",func() {
      timet:=Timestamp{time.Now()}
      caze:=CaseBuilder.SetTimestamp("LockedUntil",timet).Build()
      So(caze.LockedUntil.String(),ShouldEqual,timet.String())
    })
    Convey("builds TimestampNow fields",func() {
      //not sure how to mock this out in Go
      caze:=CaseBuilder.SetTimestampNow("LockedUntil").Build()
      So(caze.LockedUntil.String(),ShouldNotBeNil)
    })
    Convey("builds Message field",func() {
      msg:=Message{}
      caze:=CaseBuilder.SetMessage(msg).Build()
      So(caze.Message,ShouldNotBeNil)
    })
    Convey("builds LinkCollection field",func() {
      coll:=LinkCollection{}
      caze:=CaseBuilder.SetLinkCollection(coll).Build()
      So(caze.LinkCollection,ShouldNotBeNil)
    })
    Convey("builds CustomFields",func() {
      caze:=CaseBuilder.
        SetCustomField("foo1","bar1").
        SetCustomField("foo2","bar2").
        Build()
      So(caze.CustomFields["foo1"],ShouldEqual,"bar1")
      So(caze.CustomFields["foo2"],ShouldEqual,"bar2")
    })
  })
}

