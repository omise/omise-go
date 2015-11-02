package omise

import "time"

type Refund struct {
	ID          string    `json:"id"`
	Location    string    `json:"location"`
	Amount      int64     `json:"amount"`
	Currency    string    `json:"currency"`
	Charge      string    `json:"charge"`
	Transaction string    `json:"transaction"`
	Created     time.Time `json:"created"`
}

type RefundList struct {
	*List
	Data []*Refund `json:"data"`
}
