package omise

// Balance represents Omise's balance object.
// See https://www.omise.co/balance-api for more information.
type Balance struct {
	Base
	Available int64  `json:"available" pretty:""`
	Total     int64  `json:"total" pretty:""`
	Currency  string `json:"currency" pretty:""`
}
