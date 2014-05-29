package desk

import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestAllTheThings(t *testing.T) {
    fmt.Println("")
    Convey("NewLinkCollection", t, func() {
        lc := NewLinkCollection()
        Convey("should not be nil", func() {
          So(lc, ShouldNotBeNil)
        })
        Convey("should create a links map-of-maps", func() {
          So(lc.Links, ShouldNotBeNil)
          So(lc.Links, ShouldHaveSameTypeAs,make(map[string]map[string]interface{}))
        })
    })
    Convey("HasLink", t, func() {
      Convey("should be false if links is nil", func() {
        lc := NewLinkCollection()
        lc.Links=nil
        So(lc.HasLink("foo"),ShouldBeFalse)
      })
      Convey("should be false if link is not present", func(){
        lc := NewLinkCollection()
        So(lc.HasLink("foo"),ShouldBeFalse)
      })
      Convey("should be true if link is present", func(){
        lc := NewLinkCollection()
        lc.Links["foo"]=make(map[string]interface{})
        So(lc.HasLink("foo"),ShouldBeTrue)
      })
    })
    Convey("GetLinkSubItemStringValue",t,func() {
      Convey("should be blank if link is not present",func() {
        lc := NewLinkCollection()
        So(lc.GetLinkSubItemStringValue("foo","foo"),ShouldBeBlank)
      })
      Convey("should be blank if link subitem is not present",func() {
        lc := NewLinkCollection()
        lc.Links["foo"]=make(map[string]interface{})
        So(lc.GetLinkSubItemStringValue("foo","foo"),ShouldBeBlank)
      })
      Convey("should be the string value if link and subitem is present",func() {
        lc := NewLinkCollection()
        val := "bar"
        lc.Links["foo"]=make(map[string]interface{})
        lc.Links["foo"]["foo"]=val
        So(lc.GetLinkSubItemStringValue("foo","foo"),ShouldEqual,val)
      })
    })
    Convey("AddLinkSubItemStringValue",t,func() {
      Convey("when link and subitem doesn't exist", func() {
        Convey("should create new link and subitem with string value",func() {
          lc := NewLinkCollection()
          lc.AddLinkSubItemStringValue("foo","bar","blah")
          So(lc.GetLinkSubItemStringValue("foo","bar"),ShouldEqual,"blah")
        })
      })
      Convey("when link and subitem exist", func() {
        Convey("should overwrite value",func() {
          lc := NewLinkCollection()
          lc.AddLinkSubItemStringValue("foo","bar1","blah1")
          lc.AddLinkSubItemStringValue("foo","bar1","blah2")
          So(lc.GetLinkSubItemStringValue("foo","bar1"),ShouldEqual,"blah2")
        })
      })
      Convey("when link exists but subitem does not exist", func() {
        Convey("should append value to link",func() {
          lc := NewLinkCollection()
          lc.AddLinkSubItemStringValue("foo","bar1","blah1")
          lc.AddLinkSubItemStringValue("foo","bar2","blah2")
          So(lc.GetLinkSubItemStringValue("foo","bar1"),ShouldEqual,"blah1")
          So(lc.GetLinkSubItemStringValue("foo","bar2"),ShouldEqual,"blah2")
        })
      })
    })
    Convey("AddHrefLink",t,func() {
      Convey("creates correct href type link",func() {
        href := "/api/v2/customers/192220782"
        lc := NewLinkCollection()
        lc.AddHrefLink("customer",href)
        So(lc.GetLinkSubItemStringValue("customer","class"),ShouldEqual,"customer")
        So(lc.GetLinkSubItemStringValue("customer","href"),ShouldEqual,href)
      })
    })
}
