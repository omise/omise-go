package omise

type List struct {
	From string `json:"from"`
	To   string `json:"to"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`

	Location string `json:"location"`
}
