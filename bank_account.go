package omise

type BankAccount struct {
	Base
	Brand  string `json:"brand"`
	Number string `json:"number"`
	Name   string `json:"name"`

	// TODO: not documented???
	LastDigits string `json:"last_digits"`
}
