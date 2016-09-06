package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

type Search struct {
	Scope   omise.SearchScope
	Query   string
	Filters map[string]string
	Order   omise.Ordering
}

func (req *Search) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/search",
	}
}
