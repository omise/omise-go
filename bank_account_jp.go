package omise

// BankAccountJP represents Omise's bank account object for jp.
type BankAccountJP struct {
	Base
	BankCode    string          `json:"bank_code" pretty:""`
	BranchCode  string          `json:"branch_code" pretty:""`
	AccountType BankAccountType `json:"account_type" pretty:""`
	Number      string          `json:"number"`
	Name        string          `json:"name" pretty:""`
}
