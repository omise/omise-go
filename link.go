package omise

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
