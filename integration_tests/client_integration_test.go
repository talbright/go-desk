//
// Copyright (c) 2015 Highwinds Network Group, inc.
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential.
//
// @author    Scot Wells <scot.wells@highwinds.com>
// @copyright 2015 Highwinds Network Group, inc.
//
package integration_tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"strconv"
	"sync"
)

func TestClientIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	wait := sync.WaitGroup{}

	client := CreateClient()
	Convey("All requests should complete even when rate limit hit", t, func() {
		_, resp, _ := client.Case.Get("1")

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

		// wait until all the requests have completed
		wait.Wait()
		So(cases, ShouldEqual, limit + 10)
	})
}
