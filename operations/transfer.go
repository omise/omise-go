package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListTransfers struct {
	List
}

func (req *ListTransfers) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/transfers", nil}
}

type CreateTransfer struct {
	Amount    int64
	Recipient string
}

func (req *CreateTransfer) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/transfers", nil}
}

type RetreiveTransfer struct {
	TransferID string
}

func (req *RetreiveTransfer) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/transfers/" + req.TransferID, nil}
}

type UpdateTransfer struct {
	TransferID string
	Amount     int64
}

func (req *UpdateTransfer) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/transfers/" + req.TransferID, nil}
}

type DestroyTransfer struct {
	TransferID string
}

func (req *DestroyTransfer) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/transfers/" + req.TransferID, nil}
}
