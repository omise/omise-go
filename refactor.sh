#!/bin/sh

set -e

go build -o refactor ./internal/refactor
./refactor *.go internal/*.go internal/**/*.go operations/*.go schedule/*.go testdata/*.go
go fmt ./...
rm refactor
