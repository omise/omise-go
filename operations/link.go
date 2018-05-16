package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListLinks struct {
	List
}

func (req *ListLinks) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links",
		ContentType: "application/json",
	}
}

type CreateLink struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Multiple    bool   `json:"multiple"`
}

func (req *CreateLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/links",
		ContentType: "application/json",
	}
}

type RetrieveLink struct {
	LinkID string `json:"-"`
}

func (req *RetrieveLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}
