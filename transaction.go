package omise

import "time"

// Transaction represents Omise's transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Source string          `json:"source" pretty:""`
	Type   TransactionType `json:"type" pretty:""`

	Amount       int64     `json:"amount" pretty:""`
	Currency     string    `json:"currency" pretty:""`
	Transferable time.Time `json:"transferable" pretty:""`
}
