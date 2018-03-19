package internal

import (
	"net/url"
)

type Op struct {
	Endpoint Endpoint `query:"-"`
	Method   string   `query:"-"`
	Path     string   `query:"-"`

	APIKey string `query:"-"`

	Values      url.Values `query:"-"`
	Multipart   bool       `query:"-"`
	ContentType string     `query:"-"`
}

// Op implements Operation.Op and allows the struct itself be passed as an Operation
// argument directly.
func (op *Op) Op() *Op {
	return op
}

func (op *Op) KeyKind() (kind string) {
	switch {
	case op.APIKey != "":
		kind = op.APIKey
	case op.Endpoint == API:
		kind = "secret"
	case op.Endpoint == Vault:
		kind = "public"
	}
	return
}

type Operation interface {
	Op() *Op
}
