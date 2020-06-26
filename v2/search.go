package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Search represents Search object.
// See https://www.omise.co/search-api for more information.
type Search struct {
	Base
	Data []interface{} `json:"data"`
	Export string `json:"export"`
	Filters map[string]interface{} `json:"filters"`
	Location string `json:"location"`
	Order Order `json:"order"`
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Query string `json:"query"`
	Scope string `json:"scope"`
	Total int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// SearchService represents resource service.
type SearchService struct {
	*Client
}

// Search defines resource service.
func (c *Client) Search() *SearchService {
	return &SearchService{c}
}

// Search searchs
//
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
func (s *SearchService) Search(ctx context.Context, params *SearchParams) (*Search, error) {
	result := &Search{}
	err := s.Do(ctx, result, params)

	return result, err
}

// SearchParams params object.
type SearchParams struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
	Order Order `json:"order,omitempty"`
	Page int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	Query string `json:"query,omitempty"`
	Scope SearchScope `json:"scope,omitempty"`
}

// Describe describes structure of request
func (req *SearchParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/search",
		ContentType: "application/json",
	}
}

