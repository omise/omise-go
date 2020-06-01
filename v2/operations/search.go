package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

// Example:
//
//	search, search := &omise.Search{}, &Search{
//	}
//	if e := client.Do(search, search); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("search: %#v\n", search)
//
type Search struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
	Order omise.Order `json:"order,omitempty"`
	Page int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	Query string `json:"query,omitempty"`
	Scope omise.SearchScope `json:"scope,omitempty"`
}

// Describe
func (req *Search) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/search",
		ContentType: "application/json",
	}
}

