package resource

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAttachment(t *testing.T) {
	fmt.Println("")
	Convey("SetContent", t, func() {
		Convey("should read in filename contents", func() {
			attach := NewAttachment()
			err := attach.SetContent("../integration_tests/test.png")
			So(err, ShouldBeNil)
			So(attach.Content, ShouldNotBeNil)
		})
		Convey("should return an error if filename is invalid", func() {
			attach := NewAttachment()
			err := attach.SetContent("test2.png")
			So(err, ShouldNotBeNil)
		})
	})
}
