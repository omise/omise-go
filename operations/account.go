package operations

import (
	"github.com/omise/omise-go/internal"
)

type RetreiveAccount struct{}

func (req *RetreiveAccount) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/account", nil}
}
