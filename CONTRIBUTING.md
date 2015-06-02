# We love pull requests. Here's a quick guide:

1. Fork the repo.

2. Run the tests. We only take pull requests with passing tests, and it's great
to know that you have a clean slate.

3. Add a test for your change. Only refactoring and documentation changes
require no new tests. If you are adding functionality or fixing a bug, we need
a test!

4. Make the test pass.

5. Push to your fork and submit a pull request.

At this point you're waiting on us. We like to at least comment on, if not
accept, pull requests within three business days. We may suggest some changes
or improvements or alternatives.

# Tests

This project uses GoConvey for testing. There are two types of tests, unit and
integration.

## Unit Tests

To start the unit tests (from the project root):

    $ go get github.com/smartystreets/goconvey
    $ $GOPATH/bin/goconvey

Then open your browser to: http://localhost:8080

As an alternative, you may also run the following script as a shortcut, but it is 
os x specific:

    $ ./test.sh unit

## Integration Tests

Integration tests require a valid desk.com site to run against.

To start integration tests:

    $ DESK_SITE_URL=https://mysite.desk.com DESK_SITE_EMAIL=myemail@example.com DESK_SITE_PASS=mypass go test integration_tests/*.go

As an alternative, you may also run the following script as a shortcut, but it is 
os x specific:

    $ DESK_SITE_URL=https://mysite.desk.com DESK_SITE_EMAIL=myemail@example.com DESK_SITE_PASS=mypass ./test.sh integration

