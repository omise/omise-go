package internal

import (
	"net/url"
)

type Op struct {
	Endpoint Endpoint
	Method   string
	Path     string
	Values   url.Values
}

type Operation interface {
	Op() *Op
}
