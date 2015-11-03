package omise

type Balance struct {
	Base
	Available int64  `json:"available"`
	Total     int64  `json:"total"`
	Currency  string `json:"currency"`
}
