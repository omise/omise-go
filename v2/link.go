package omise

import (
	"time"
)

// Link represents Links object.
// See https://www.omise.co/links-api for more information.
type Link struct {
	Base
	Amount int `json:"amount"`
	Charges *ChargeList `json:"charges"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Location string `json:"location"`
	Multiple bool `json:"multiple"`
	PaymentURI string `json:"payment_uri"`
	Title string `json:"title"`
	Used bool `json:"used"`
	UsedAt time.Time `json:"used_at"`
}
