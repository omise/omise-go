package omise

// BankAccount represents Omise's bank account object.
// See https://www.omise.co/bank-account-api for more information.
type BankAccount struct {
	Base
	Brand      string `json:"brand" pretty:""`
	Number     string `json:"number"`
	LastDigits string `json:"last_digits" pretty:""`
	Name       string `json:"name" pretty:""`
}
