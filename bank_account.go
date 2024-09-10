package omise

// BankAccount represents Omise's bank account object.
// See https://www.omise.co/bank-account-api for more information.
type BankAccountRequest struct {
	Base
	Brand      string `json:"brand"`
	LastDigits string `json:"last_digits"`
	Name       string `json:"name"`

	// for Omise Japan
	BankCode   string          `json:"bank_code"`
	BranchCode string          `json:"branch_code"`
	Type       BankAccountType `json:"type"`
	Number     string          `json:"number"`
}

type BankAccount struct {
	Base
	Brand      string `json:"brand"`
	LastDigits string `json:"last_digits"`
	Name       string `json:"name"`

	// for Omise Japan
	BankCode      string          `json:"bank_code"`
	BranchCode    string          `json:"branch_code"`
	Type          BankAccountType `json:"type"`
	AccountNumber string          `json:"account_number"`
}

// BankAccountType BankAccount an enumeration of possible types of BackAccount(s) which can be
type BankAccountType string

// BankAccountType can be one of the following list of constants:
const (
	Normal  BankAccountType = "normal"
	Current BankAccountType = "current"
)
