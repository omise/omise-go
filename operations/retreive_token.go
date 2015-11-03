package operations

import (
	"github.com/omise/omise-go/internal"
)

type RetreiveToken struct {
	ID string
}

func (token *RetreiveToken) Endpoint() (internal.Endpoint, string, string) {
	return internal.Vault, "GET", "/tokens/" + token.ID
}

func (*RetreiveToken) Payload() (interface{}, error) {
	return nil, nil
}
