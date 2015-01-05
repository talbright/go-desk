package desk

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResourcePath(t *testing.T) {
	fmt.Println("")
	Convey("Path", t, func() {
		Convey("should generate a valid member path", func() {
			path := NewIdentityResourcePath("1",NewCase())
			So(path.Path(), ShouldEqual, "cases/1")
		})
		Convey("should generate a valid collection path", func() {
			path := NewIdentityResourcePath("1",NewCase()).SetCollection()
			So(path.Path(), ShouldEqual, "cases")
		})
		Convey("should inclue a prefix in the path", func() {
			path := NewIdentityResourcePath("1",NewCase()).SetPrefix("api/v2")
			So(path.Path(), ShouldEqual, "api/v2/cases/1")
		})
		Convey("should include a suffix in the path", func() {
			path := NewIdentityResourcePath("1",NewCase()).SetSuffix("replies/1")
			So(path.Path(), ShouldEqual, "cases/1/replies/1")
		})
		Convey("should include an action in the path", func() {
			path := NewIdentityResourcePath("1",NewCase()).SetAction("preview")
			So(path.Path(), ShouldEqual, "cases/1/preview")
		})
		Convey("should include a nested resource in the path", func() {
			draft := NewDraft()
			draft.SetResourceId("2")
			caze := NewCase()
			caze.SetResourceId("1")
			path := caze.GetResourcePath(caze)
			path.Nested = draft
			So(path.Path(), ShouldEqual, "cases/1/draft/2")
		})
		Convey("should append additional paths", func() {
			path1 := NewIdentityResourcePath("1",NewCase())
			path2 := NewIdentityResourcePath("2",NewCase())
			path1.AppendPath(path2)
			So(path1.Path(), ShouldEqual, "cases/1/cases/2")
		})
		Convey("should create case replies path", func() {
			path := NewIdentityResourcePath("1",NewCase()).SetAction("replies").SetNested(NewDraft())
			So(path.Path(), ShouldEqual, "cases/1/replies/draft")
		})
	})
}
