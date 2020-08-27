package omise

// RecipientFailureCode represents an enumeration of possible values for Recipient.
type RecipientFailureCode string

// RecipientFailureCode can be one of the following list of constants:
const (
	RecipientAccountNotFound RecipientFailureCode = "account_not_found"
	RecipientBankNotFound RecipientFailureCode = "bank_not_found"
	RecipientNameMismatch RecipientFailureCode = "name_mismatch"
)

