package omise

// Link represents Omise's link object.
// See https://www.omise.co/links-api for more information.
type Link struct {
	Base
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
	Used       bool   `json:"used"`
	Multiple   bool   `json:"multiple"`
	PaymentURI string `json:"payment_uri"`

	Title       string     `json:"title"`
	Description string     `json:"description"`
	Charges     ChargeList `json:"charges"`
}
