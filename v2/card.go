package omise

// Card represents Card object.
// See https://www.omise.co/cards-api for more information.
type Card struct {
	Base
	Bank string `json:"bank"`
	Brand string `json:"brand"`
	City *string `json:"city"`
	Country string `json:"country"`
	ExpirationMonth int `json:"expiration_month"`
	ExpirationYear int `json:"expiration_year"`
	Financing string `json:"financing"`
	Fingerprint string `json:"fingerprint"`
	FirstDigits *string `json:"first_digits"`
	LastDigits string `json:"last_digits"`
	Location string `json:"location"`
	Name string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	PostalCode *string `json:"postal_code"`
	SecurityCodeCheck bool `json:"security_code_check"`
	State *string `json:"state"`
	Street1 *string `json:"street1"`
	Street2 *string `json:"street2"`
}
