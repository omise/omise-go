package omise

//go:generate go run internal/generator/main.go list_types
//go:generate go fmt list_types.go

// List structure contains fields that are common to list objects returned by the Omise's
// REST API. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type List struct {
	Base
	From string `json:"from"`
	To   string `json:"to"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`

	Order Ordering `json:"order"`
}
