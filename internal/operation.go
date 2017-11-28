package internal

import (
	"net/url"
)

type Op struct {
	Endpoint Endpoint `query:"-"`
	Method   string   `query:"-"`
	Path     string   `query:"-"`

	ForceUsePublicKey bool `query:"-"`

	Values      url.Values `query:"-"`
	Multipart   bool       `query:"-"`
	ContentType string     `query:"-"`
}

// Op implements Operation.Op and allows the struct itself be passed as an Operation
// argument directly.
func (op *Op) Op() *Op {
	return op
}

type Operation interface {
	Op() *Op
}
