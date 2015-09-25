//
// Copyright (c) 2015 Highwinds Network Group, inc.
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential.
//
// @author    Scot Wells <scot.wells@highwinds.com>
// @copyright 2015 Highwinds Network Group, inc.
//
package types_test
import (
	"testing"
	"github.com/talbright/go-desk/types"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

func TestFilter(t *testing.T) {
	filter := types.NewInsightsV3Filter()

	Convey("TestFilter", t, func() {
		Convey("test correct json output", func() {
			filter.Type = types.TYPE_EXCLUDE
			filter.Field = types.FIELD_LABELS

			json, err := json.Marshal(filter)
		})
	})
}

func TestFilterGroup(t *testing.T) {
	group := types.NewInsightsV3FilterGroup()

	Convey("TestFilterGroup", t, func() {
		Convey("Test add filter", func() {
			So(len(group.Filters), ShouldEqual, 0)
			filter := types.NewInsightsV3Filter()
			filter.Field = "Custom Fields"
			filter.Type = "Include"

			group.Add(filter)
			So(len(group.Filters), ShouldEqual, 1)

			j, err := group.MarshalJSON()

			So(err, ShouldBeNil)
			So(j, ShouldNotBeNil)

			json := string(j)
			So(json, ShouldContainSubstring, "Custom Fields")
			So(json, ShouldContainSubstring, "Include")
		})
	})
}
