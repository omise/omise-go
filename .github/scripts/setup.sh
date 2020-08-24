#!/bin/sh

go get github.com/alecthomas/gometalinter && gometalinter --install
go generate . ./operations
go install . ./operations