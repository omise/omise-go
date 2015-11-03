package omise

type Refund struct {
	Base
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Charge      string `json:"charge"`
	Transaction string `json:"transaction"`
}

type RefundList struct {
	*List
	Data []*Refund `json:"data"`
}
