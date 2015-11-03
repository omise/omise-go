package operations

import (
	"github.com/omise/omise-go/internal"
)

type RetreiveBalance struct{}

func (req *RetreiveBalance) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/balance", nil}
}
