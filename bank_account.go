package omise

// BankAccount represents Omise's bank account object.
// See https://www.omise.co/bank-account-api for more information.
type BankAccount struct {
	Base
	Brand      string `json:"brand" pretty:""`
	Number     string `json:"number"`
	LastDigits string `json:"last_digits" pretty:""`
	Name       string `json:"name" pretty:""`

	// for Omise Japan
	BankCode    string          `json:"bank_code" pretty:""`
	BranchCode  string          `json:"branch_code" pretty:""`
	AccountType BankAccountType `json:"account_type" pretty:""`
}

// BankAccountType BankAccount an enumeration of possible types of BackAccount(s) which can be
type BankAccountType string

// BankAccountType can be one of the following list of constants:
const (
	Normal  BankAccountType = "normal"
	Current BankAccountType = "current"
)
