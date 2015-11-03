package omise

type Dispute struct {
	Base
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	Charge   string `json:"charge"`
}
