package omise

// BankAccountType BankAccount an enumeration of possible types of BackAccount(s) which can be
type BankAccountType string

// BankAccountType can be one of the following list of constants:
const (
	Normal  BankAccountType = "normal"
	Current BankAccountType = "current"
)
