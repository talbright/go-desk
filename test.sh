#!/bin/bash

open http://localhost:8080
$GOPATH/bin/goconvey -short=true -depth=1
