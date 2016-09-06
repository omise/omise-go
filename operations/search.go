package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	result, search := &omise.ChargeSearchResult{}, &Search{
//		Scope: omise.ChargeScope,
//		Query: "chrg_test_12345",
//	}
//	if e := client.Do(result, search); e != nil {
//		panic(e)
//	}
//
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
