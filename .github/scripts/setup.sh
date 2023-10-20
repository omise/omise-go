#!/bin/sh

go generate . ./operations
go install . ./operations
