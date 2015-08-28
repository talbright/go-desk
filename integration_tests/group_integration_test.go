package integration_tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestGroupIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a group by ID", t, func() {
		group, _, err := client.Group.Get(fmt.Sprintf("%d", DefaultGroupId))
		So(err, ShouldBeNil)
		So(group, ShouldNotBeNil)
	})

	Convey("should be able to retrieve a list of groups", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.Group.List(&listParams)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to retrieve a list of users for a group", t, func() {
		collection, _, err := client.Group.Users(fmt.Sprintf("%d", DefaultGroupId))
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})
}
