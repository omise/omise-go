package omise

// TransactionDirection represents an enumeration of possible values for Transaction.
type TransactionDirection string

// TransactionDirection can be one of the following list of constants:
const (
	TransactionCredit TransactionDirection = "credit"
	TransactionDebit TransactionDirection = "debit"
)

