#!/bin/sh

TMPFILE=`mktemp -t omise-go-XXXXXX`
GOPATH=`go env GOPATH 2>/dev/null`

echo "using TMPFILE=$TMPFILE"
echo "using GOPATH=$GOPATH"

check() {
  printf "checking for `basename $1` ... "
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

perform() {
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

check go                       "needs go from http://golang.org"
check $GOPATH/bin/gometalinter "needs gometalinter from https://github.com/alecthomas/gometalinter"

perform linters    gometalinter --fast --aggregate
perform tests      go test ./...  -coverprofile cover.out

echo "\x1B[38;5;2msuccess.\x1B[0m"
