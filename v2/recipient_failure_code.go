package omise

type RecipientFailureCode string

const (
	RecipientAccountNotFound RecipientFailureCode = "account_not_found"
	RecipientBankNotFound RecipientFailureCode = "bank_not_found"
	RecipientNameMismatch RecipientFailureCode = "name_mismatch"
)
