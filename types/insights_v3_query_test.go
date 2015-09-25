package types_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/talbright/go-desk/types"
	"testing"
)

func TestFields(t *testing.T) {
	query := types.NewInsightsV3Query()

	Convey("SetFieldsCorrectly", t, func() {
		Convey("Should set single field correctly", func() {
			query.AddField("testField")

			// we've only added a single field
			So(len(query.Fields), ShouldEqual, 1)
		})

		Convey("Should set multiple fields correctly", func() {
			query.AddFields([]string{"test1", "test2"})

			// added a single field before, now adding another two.. 1 + 2 = 3 :)
			So(len(query.Fields), ShouldEqual, 3)
		})

		Convey("Shouldn't allow duplicates", func() {
			query.AddField("test2")

			So(len(query.Fields), ShouldEqual, 3)
		})
		Convey("Should convert to JSON properly", func() {
			json := query.String()

			// test that the json got generated correctly
			So(json, ShouldContainSubstring, `"fields":["testField","test1","test2"]`)
		})
	})
}

func TestTimeSet(t *testing.T) {
	query := types.NewInsightsV3Query()

	Convey("SetTimeCorrectly", t, func() {
		Convey("Be able to set the time", func() {
			time := types.NewInsightsV3Time()

			time.WindowSize = "hour"

			query.SetReportTime(time)

			So(query.Time.WindowSize, ShouldEqual, "hour")
		})
	})
}

func TestDimensions(t *testing.T) {
	query := types.NewInsightsV3Query()

	Convey("WorkWithDimensions", t, func() {
		Convey("Be able to set Dimension1", func() {
			query.SetDimension1("action_agent")

			So(query.Dimension1, ShouldEqual, "action_agent")
		})

		Convey("Should be able to set Dimension1Value", func() {
			query.SetDimension1Values("test")

			So(query.Dimension1Values, ShouldEqual, "test")
		})

		Convey("Should be able to set Dimension2", func() {
			query.SetDimension2("test2")

			So(query.Dimension2, ShouldEqual, "test2")
		})

		Convey("Should be able to set Dimension2Values", func() {
			query.SetDimension2Values("test3")

			So(query.Dimension2Values, ShouldEqual, "test3")
		})

		Convey("Should convert to JSON correctly", func() {
			json := query.String()

			So(json, ShouldContainSubstring, `"dimension1":"action_agent"`)
			So(json, ShouldContainSubstring, `"dimension1_values":"test"`)
			So(json, ShouldContainSubstring, `"dimension2":"test2"`)
			So(json, ShouldContainSubstring, `"dimension2_values":"test3"`)
		})
	})
}

func TestFilters(t *testing.T) {
	query := types.NewInsightsV3Query()

	Convey("QueryFilters", t, func() {
		Convey("Add a filter group", func() {
			query.AddFilters(types.NewInsightsV3FilterGroup())
			So(len(query.Filters), ShouldEqual, 1)

			query.AddFilters(types.NewInsightsV3FilterGroup())
			So(len(query.Filters), ShouldEqual, 2)
		})
	})
}
