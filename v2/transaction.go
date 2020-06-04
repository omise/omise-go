package omise

import (
	"time"
)

// Transaction represents Transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	Direction *TransactionDirection `json:"direction"`
	Key string `json:"key"`
	Location string `json:"location"`
	Origin string `json:"origin"`
	TransferableAt time.Time `json:"transferable_at"`
}