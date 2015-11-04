#!/usr/bin/make

gen:
	go install ./internal/generator
	$(GOPATH)/bin/generator

deps:
	go get -v -u ./...

test:
	go test -v ./...
