package integration_tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	. "github.com/talbright/go-desk/types"
	"time"
)

func TestInsightsV3Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipping in short mode")
	}
	client := CreateClient()

	Convey("should be able to run report", t, func() {
		query := NewInsightsV3Query()
		query.Dimension1 = "action_agent"
		query.Fields = []string{"case_reopens", "resolved_cases", "inbound_interactions"}
		// view 10 week period
		query.Time.Min = Timestamp{time.Now().Add(-10 * 7 * 24 * time.Hour).UTC()}
		query.Time.Max = Timestamp{time.Now().UTC()}
		query.Time.WindowSize = "hour"

		report, resp, err := client.InsightsV3.Report(query)

		data := report.GetData()

		// can only look at the results if we got a response
		if len(data) > 0 {
			r := data[0]

			// check that the values were returned
			for _, field := range report.Request.Fields {
				if _, ok := r[field]; !ok {
					t.Error("field ", field, " was not found in response")
				}
			}
		}

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

