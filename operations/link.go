package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListLinks struct {
	List
}

func (req *ListLinks) Describe() *internal.Description {
	return &internal.Description{
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

func (req *CreateLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/links",
	}
}

type RetrieveLink struct {
	LinkID string `query:"-"`
}

func (req *RetrieveLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/links/" + req.LinkID,
	}
}
