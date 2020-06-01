package omise

// BankAccount represents Bank Account object.
// See https://www.omise.co/bank-account-api for more information.
type BankAccount struct {
	Base
	BankCode string `json:"bank_code"`
	BranchCode string `json:"branch_code"`
	Brand string `json:"brand"`
	LastDigits string `json:"last_digits"`
	Name string `json:"name"`
	Type BankAccountType `json:"type"`
}

