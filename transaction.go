package omise

import "time"

// Transaction represents Omise's transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Source string          `json:"source"`
	Type   TransactionType `json:"type"`

	Amount       int64     `json:"amount"`
	Currency     string    `json:"currency"`
	Transferable time.Time `json:"transferable"`
}
