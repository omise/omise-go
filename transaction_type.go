package omise

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit                  = "debit"
)
