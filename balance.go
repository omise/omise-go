package omise

// Balance represents Omise's balance object.
// See https://www.omise.co/balance-api for more information.
type Balance struct {
	Base
	Available int64  `json:"available"`
	Total     int64  `json:"total"`
	Currency  string `json:"currency"`
}
