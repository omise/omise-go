package operations

import (
	"github.com/omise/omise-go/internal"
)

type RetreiveBalance struct{}

func (*RetreiveBalance) Endpoint() (internal.Endpoint, string, string) {
	return internal.API, "GET", "/balance"
}
