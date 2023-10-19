#!/bin/sh

go get -u github.com/alecthomas/gometalinter && gometalinter --install
go generate . ./operations
go install . ./operations
