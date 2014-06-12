package desk

import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestHal(t *testing.T) {
    fmt.Println("")
    Convey("GetStringId", t, func() {
        Convey("should return stringified Id", func() {
          hal := NewHal()
          hal.Id = new(int)
          *hal.Id = 123
          So(hal.GetStringId(),ShouldEqual,"123")
        })
    })
    Convey("GetId", t, func() {
        Convey("should use Id field if present", func() {
          hal := NewHal()
          hal.Id = new(int)
          *hal.Id = 123
          So(hal.GetId(),ShouldEqual,123)
        })
        Convey("should use href field if Id field is not present", func() {
          hal := NewHal()
          hal.AddHrefLink("self","/api/v2/cases/75/replies/123")
          So(hal.GetId(),ShouldEqual,123)
        })
    })
    Convey("NewHal", t, func() {
        hal := NewHal()
        Convey("should not be nil", func() {
          So(hal, ShouldNotBeNil)
        })
        Convey("should create a links map-of-maps", func() {
          So(hal.Links, ShouldNotBeNil)
          So(hal.Links, ShouldHaveSameTypeAs,make(map[string]map[string]interface{}))
        })
    })
    Convey("HasLink", t, func() {
      Convey("should be false if links is nil", func() {
        hal := NewHal()
        hal.Links=nil
        So(hal.HasLink("foo"),ShouldBeFalse)
      })
      Convey("should be false if link is not present", func(){
        hal := NewHal()
        So(hal.HasLink("foo"),ShouldBeFalse)
      })
      Convey("should be true if link is present", func(){
        hal := NewHal()
        hal.Links["foo"]=make(map[string]interface{})
        So(hal.HasLink("foo"),ShouldBeTrue)
      })
    })
    Convey("GetLinkSubItemStringValue",t,func() {
      Convey("should be blank if link is not present",func() {
        hal := NewHal()
        So(hal.GetLinkSubItemStringValue("foo","foo"),ShouldBeBlank)
      })
      Convey("should be blank if link subitem is not present",func() {
        hal := NewHal()
        hal.Links["foo"]=make(map[string]interface{})
        So(hal.GetLinkSubItemStringValue("foo","foo"),ShouldBeBlank)
      })
      Convey("should be the string value if link and subitem is present",func() {
        hal := NewHal()
        val := "bar"
        hal.Links["foo"]=make(map[string]interface{})
        hal.Links["foo"]["foo"]=val
        So(hal.GetLinkSubItemStringValue("foo","foo"),ShouldEqual,val)
      })
    })
    Convey("AddLinkSubItemStringValue",t,func() {
      Convey("when link and subitem doesn't exist", func() {
        Convey("should create new link and subitem with string value",func() {
          hal := NewHal()
          hal.AddLinkSubItemStringValue("foo","bar","blah")
          So(hal.GetLinkSubItemStringValue("foo","bar"),ShouldEqual,"blah")
        })
      })
      Convey("when link and subitem exist", func() {
        Convey("should overwrite value",func() {
          hal := NewHal()
          hal.AddLinkSubItemStringValue("foo","bar1","blah1")
          hal.AddLinkSubItemStringValue("foo","bar1","blah2")
          So(hal.GetLinkSubItemStringValue("foo","bar1"),ShouldEqual,"blah2")
        })
      })
      Convey("when link exists but subitem does not exist", func() {
        Convey("should append value to link",func() {
          hal := NewHal()
          hal.AddLinkSubItemStringValue("foo","bar1","blah1")
          hal.AddLinkSubItemStringValue("foo","bar2","blah2")
          So(hal.GetLinkSubItemStringValue("foo","bar1"),ShouldEqual,"blah1")
          So(hal.GetLinkSubItemStringValue("foo","bar2"),ShouldEqual,"blah2")
        })
      })
    })
    Convey("AddHrefLink",t,func() {
      Convey("creates correct href type link",func() {
        href := "/api/v2/customers/192220782"
        hal := NewHal()
        hal.AddHrefLink("customer",href)
        So(hal.GetLinkSubItemStringValue("customer","class"),ShouldEqual,"customer")
        So(hal.GetLinkSubItemStringValue("customer","href"),ShouldEqual,href)
      })
    })
}
