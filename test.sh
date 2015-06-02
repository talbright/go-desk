#!/bin/bash

function run_unit_tests {
	open http://localhost:8080
	$GOPATH/bin/goconvey -short=true
}

function run_integration_tests {
	go test -v integration_tests/*.go
}

case $1 in
	unit)
		run_unit_tests
		;;
	integration)
		run_integration_tests
		;;
	*)
		unit_tests
esac
