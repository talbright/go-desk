package integration_tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration tests are skipped in short mode.")
	}
	client := CreateClient()

	Convey("should be able to retrieve a user by ID", t, func() {
		user, _, err := client.User.Get(fmt.Sprintf("%d", DefaultUserId))
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
	})

	Convey("should be able to retrieve a list of users", t, func() {
		listParams := url.Values{}
		listParams.Add("sort_field", "created_at")
		listParams.Add("sort_direction", "asc")
		collection, _, err := client.User.List(&listParams)
		So(err, ShouldBeNil)
		So(*collection.TotalEntries, ShouldBeGreaterThan, 0)
		So(*collection.Embedded, ShouldNotBeNil)
	})

	Convey("should be able to update a user", t, func() {
		//TODO Only the currently logged-in user may be updated.
	})

}
