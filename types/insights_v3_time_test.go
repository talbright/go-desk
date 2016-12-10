package types_test

import (
	"testing"
	"github.com/talbright/go-desk/types"
	."github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

func TestTime(t *testing.T) {
	time := types.NewInsightsV3Time()

	Convey("WindowSize", t, func() {
		Convey("Should set correctly", func() {
			err := time.SetWindowSize("none")
			So(err, ShouldBeNil)

			err = time.SetWindowSize("invalid")
			So(err, ShouldNotBeNil)
		})

		Convey("should convert to Json correctly", func() {
			jsonBytes, err := json.Marshal(time)

			json := string(jsonBytes)

			So(err, ShouldBeNil)
			So(json, ShouldContainSubstring, "min")
			So(json, ShouldContainSubstring, "max")
			So(json, ShouldContainSubstring, `"window_size":"none"`)
		})
	})

}
