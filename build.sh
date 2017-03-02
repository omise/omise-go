#!/bin/sh

TMPFILE=`mktemp -t omise-go-XXXXXX`
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

  $@ 2>&1 > $TMPFILE
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
check go-bindata   "needs go-bindata from https://github.com/jteeuwen/go-bindata"

perform linters    gometalinter --fast -e "credentials,HIGH,LOW"
perform generators go generate ./...
perform builds     go install ./...
perform tests      go test ./...

echo "\x1B[38;5;2msuccess.\x1B[0m"
