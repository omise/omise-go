package omise

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

