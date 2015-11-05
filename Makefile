#!/usr/bin/make

PKG := github.com/omise/omise-go

GO    := go
GODOC := godoc

gen:
	@$(GO) get -v ./internal/generator
	@$(GOPATH)/bin/generator

deps:
	@$(GO) get -v -u ./...

test:
	@$(GO) test -v ./...

doc:
	@echo starting godoc on port 9090
	@echo package doc is at http://0.0.0.0:9090/pkg/github.com/omise/omise-go/
	@$(GODOC) -v -http=":9090"

