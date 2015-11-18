package omise

// Transaction represents Omise's transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Amount   int64           `json:"amount" pretty:""`
	Currency string          `json:"currency" pretty:""`
	Type     TransactionType `json:"type" pretty:""`
}
