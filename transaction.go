package omise

type Transaction struct {
	Base
	Amount   int64           `json:"amount"`
	Currency string          `json:"currency"`
	Type     TransactionType `json:"type"`
}
