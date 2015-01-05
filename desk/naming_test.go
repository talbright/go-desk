package desk

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNaming(t *testing.T) {
	fmt.Println("")
	Convey("GetResourceName", t, func() {
		Convey("should return correct plural name by default", func() {
			draft := &Draft{}
			naming := &Naming{}
			naming.SetResourceName(draft)
			So(naming.GetResourceName(), ShouldEqual, "drafts")
		})
		Convey("should singularize name", func() {
			draft := &Draft{}
			naming := &Naming{}
			naming.SetResourceName(draft)
			naming.Singularize()
			So(naming.GetResourceName(), ShouldEqual, "draft")
		})
		Convey("should handle nil correctly", func() {
			naming := &Naming{}
			So(naming.GetResourceName(), ShouldEqual, "")
		})
	})
}
