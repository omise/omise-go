package omise

// RecipientType represents an enumeration of possible types of Transaction which can be
// one of the following list of constants:
type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)
