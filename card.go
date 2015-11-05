package omise

import (
	"time"
)

// Card represents Omise's card object.
// See https://www.omise.co/cards-api for more information.
type Card struct {
	Base
	Country    string `json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Financing  string `json:"financing"`
	LastDigits string `json:"last_digits"`
	Brand      string `json:"brand"`

	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`

	Fingerprint       string    `json:"fingerprint"`
	Name              string    `json:"name"`
	Created           time.Time `json:"created"`
	SecurityCodeCheck bool      `json:"security_code_check"`
}
