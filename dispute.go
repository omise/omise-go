package omise

type Dispute struct {
	Base
	Amount   int64         `json:"amount"`
	Currency string        `json:"currency"`
	Status   DisputeStatus `json:"status"`
	Message  string        `json:"message"`
	Charge   string        `json:"charge"`
}

type DisputeList struct {
	List
	Data []*Dispute `json:"data"`
}
