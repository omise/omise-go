#!/bin/sh
GOPATH=`go env GOPATH 2>/dev/null`
go test -v ./...  -coverprofile cover.out
