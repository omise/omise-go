package internal

type Description struct {
	Endpoint Endpoint
	Method   string
	Path     string

	APIKey string

	ContentType string
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
