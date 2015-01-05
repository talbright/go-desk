package desk

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDraft(t *testing.T) {
	fmt.Println("")
	Convey("GetResourceName", t, func() {
		Convey("should return correct name", func() {
			draft := NewDraft()
			So(draft.GetResourceName(), ShouldEqual, "draft")
		})
	})
}

