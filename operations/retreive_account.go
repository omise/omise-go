package operations

import (
	"github.com/omise/omise-go/internal"
)

type RetreiveAccount struct{}

func (*Retreiveaccount) Endpoint() (internal.Endpoint, string, string) {
	return internal.API, "GET", "/account"
}
