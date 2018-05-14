package internal

import "net/url"

type Description struct {
	Endpoint Endpoint `query:"-"`
	Method   string   `query:"-"`
	Path     string   `query:"-"`

	APIKey string `query:"-"`

	Values      url.Values `query:"-"`
	ContentType string     `query:"-"`
}

// Op implements Operation.Describe and allows the struct itself be passed as an Operation
// argument directly.
func (op *Description) Describe() *Description {
	return op
}

func (op *Description) KeyKind() (kind string) {
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
