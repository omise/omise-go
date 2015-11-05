#!/usr/bin/make

GO := go

gen:
	@$(GO) get -v ./internal/generator
	@$(GOPATH)/bin/generator

deps:
	@$(GO) get -v -u ./...

test:
	@$(GO) test -v ./...
