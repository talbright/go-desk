#!/bin/bash

GOPATH=`godep path`:$GOPATH

function run_unit_tests {
	open http://localhost:8080
	goconvey -short=true
}

function run_integration_tests {
	godep go test -v integration_tests/*.go
}

case $1 in
	unit)
		run_unit_tests
		;;
	integration)
		run_integration_tests
		;;
	*)
		run_unit_tests
esac
