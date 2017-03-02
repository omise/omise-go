package omise

// TransactionType represents an enumeration of possible types of Transaction.
type TransactionType string

// TransactionType can be one of the following list of constants:
const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)
