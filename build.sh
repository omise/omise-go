#!/bin/sh

TMPFILE=`mktemp -t omise-go`
GOPATH=`go env GOPATH 2>/dev/null`

echo "using TMPFILE=$TMPFILE"
echo "using GOPATH=$GOPATH"

function check() {
  printf "checking for $1 ... "
  command -v $1 2>&1 >/dev/null
  if [ ! $? -eq 0 ]
  then
    echo "not ok."
    echo "$@"
    echo "\x1B[38;5;9mfailed.\x1B[0m"
    exit 1
  fi

  echo "ok."
}

function perform() {
  printf "running $1 ... "
  shift

  $@ > $TMPFILE
  if [ ! $? -eq 0 ]
  then
    echo "not ok."
    cat $TMPFILE
    echo "\x1B[38;5;9mfailed.\x1B[0m"
    exit 1
  fi

  echo "ok."
}


check go           "needs go from http://golang.org"
check gometalinter "needs gometalinter from https://github.com/alecthomas/gometalinter"

perform linters    gometalinter --fast
perform generators go generate ./...
perform tests      go test ./...
perform builds     go install ./...

echo "\x1B[38;5;2msuccess.\x1B[0m"
