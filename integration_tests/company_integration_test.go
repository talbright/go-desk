package integration_tests

import (
	// "log"
	"testing"
	// "fmt"
	"net/url"
	// "time"
	. "github.com/smartystreets/goconvey/convey"
	// resource "github.com/talbright/go-desk/resource"
	// types "github.com/talbright/go-desk/types"
)

func TestCompanyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a list of companies", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Company.List(&listParams)
		So(err,ShouldBeNil)
		So(*collection.TotalEntries,ShouldBeGreaterThan,0)
		So(*collection.Embedded,ShouldNotBeNil)
	})
}

