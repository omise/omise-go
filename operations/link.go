package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListLinks struct {
	List
}

func (req *ListLinks) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/links",
	}
}

type CreateLink struct {
	Amount      int64
	Currency    string
	Title       string
	Description string
	Multiple    bool
}

func (req *CreateLink) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/links",
	}
}

type RetrieveLink struct {
	LinkID string `query:"-"`
}

func (req *RetrieveLink) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/links/" + req.LinkID,
	}
}
