package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListTransaction struct {
	List
}

func (req *ListTransaction) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/transactions", nil}
}

type RetrieveTransaction struct {
	TransactionID string `query:"-"`
}

func (req *RetrieveTransaction) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/transactions/" + req.TransactionID, nil}
}
