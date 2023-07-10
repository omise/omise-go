#!/bin/sh
GOPATH=`go env GOPATH 2>/dev/null`
go test ./...  -coverprofile cover.out
