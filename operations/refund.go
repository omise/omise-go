package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListRefunds struct {
	ChargeID string `query:"-"`
	List
}

func (req *ListRefunds) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID + "/refunds", nil}
}

type CreateRefund struct {
	ChargeID string `query:"-"`
	Amount   int64
}

func (req *CreateRefund) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/charges/" + req.ChargeID + "/refunds", nil}
}

type RetreiveRefund struct {
	ChargeID string `query:"-"`
	RefundID string `query:"-"`
}

func (req *RetreiveRefund) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID + "/refunds/" + req.RefundID, nil}
}
