package omise

type List struct {
	Base
	From string `json:"from"`
	To   string `json:"to"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}
