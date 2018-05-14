package omise

//go:generate go run internal/generator/main.go search_result_types
//go:generate go fmt search_result_types.go

// SearchResult structure contains fields that are common to search result objects
// returned by Omise's REST API.
type SearchResult struct {
	Base
	Scope   SearchScope       `json:"scope"`
	Query   string            `json:"query"`
	Filters map[string]string `json:"filters"`

	Page       int `json:"page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`

	Order Ordering `json:"order"`
}
