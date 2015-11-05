package omise

// Transaction represents Omise's transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Amount   int64           `json:"amount"`
	Currency string          `json:"currency"`
	Type     TransactionType `json:"type"`
}
