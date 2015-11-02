package omise

import (
	"time"
)

type Card struct {
	ID         string  `json:"id"`
	Live       bool    `json:"livemode"`
	Location   *string `json:"location"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	PostalCode string  `json:"postal_code"`
	Financing  string  `json:"financing"`
	LastDigits string  `json:"last_digits"`
	Brand      string  `json:"brand"`

	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`

	Fingerprint string    `json:"fingerprint"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`

	SecurityCodeCheck bool `json:"security_code_check"`
}
