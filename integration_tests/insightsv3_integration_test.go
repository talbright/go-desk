package integration_tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	. "github.com/talbright/go-desk/types"
	"time"
	"strconv"
	"sync"
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

		// can only look at the results if we got a response
		if len(report.Data) > 0 {
			r := report.Data[0]

			// check that the values were returned
			for _, field := range report.Request.Fields {
				if _, ok := r.Data[field]; !ok {
					t.Error("field ", field, " was not found in response")
				}
			}
		}

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})

	Convey("report should run when rate-limited", t, func() {
		// exhaust the rate limit
		_, resp, _ := client.Case.Get("1")
		wait := sync.WaitGroup{}
		// get the remaining amount of requests that are available from the API
		ratelimit := resp.Header.Get("X-Rate-Limit-Remaining")
		limit, _ := strconv.Atoi(ratelimit)
		cases := 0

		// send the number of requests remaining plus 10, seems like sometimes the
		// Desk API will let you pass the rate limit threshold
		for i := 0; i < limit + 10; i++ {
			wait.Add(1)
			go func(){
				// issue a request against the API
				// this request will block when the api limit has been hit
				_, resp, _ := client.Case.Get("1")

				// test that we had a successful response
				if resp.StatusCode == 200 {
					cases++
				}

				wait.Done()
			}()
		}

		query := NewInsightsV3Query()
		query.Dimension1 = "action_agent"
		query.Fields = []string{"case_reopens", "resolved_cases", "inbound_interactions"}
		// view 10 week period
		query.Time.Min = Timestamp{time.Now().Add(-10 * 7 * 24 * time.Hour).UTC()}
		query.Time.Max = Timestamp{time.Now().UTC()}
		query.Time.WindowSize = "hour"

		// at this point the request should be rate-limited, but should still complete
		report, resp, err := client.InsightsV3.Report(query)

		// can only look at the results if we got a response
		if len(report.Data) > 0 {
			r := report.Data[0]

			// check that the values were returned
			for _, field := range report.Request.Fields {
				if _, ok := r.Data[field]; !ok {
					t.Error("field ", field, " was not found in response")
				}
			}
		}

		// wait until all the requests have completed
		wait.Wait()
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})

	Convey("should handle two report requests", t, func(){
		wait := sync.WaitGroup{}

		query := NewInsightsV3Query()
		query.Dimension1 = "action_agent"
		query.Fields = []string{"case_reopens", "resolved_cases", "inbound_interactions"}
		// view 10 week period
		query.Time.Min = Timestamp{time.Now().Add(-10 * 7 * 24 * time.Hour).UTC()}
		query.Time.Max = Timestamp{time.Now().UTC()}
		query.Time.WindowSize = "hour"

		// send the first request
		client.InsightsV3.Report(query)
		// only worry about the second request
		report, resp, err := client.InsightsV3.Report(query)


		// can only look at the results if we got a response
		if len(report.Data) > 0 {
			r := report.Data[0]

			// check that the values were returned
			for _, field := range report.Request.Fields {
				if _, ok := r.Data[field]; !ok {
					t.Error("field ", field, " was not found in response")
				}
			}
		}

		// wait until all the requests have completed
		wait.Wait()
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}

