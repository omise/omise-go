package omise

import (
	"time"
)

type Charge struct {
	ID          string `json:"id"`
	Live        bool   `json:"livemode"`
	Location    string `json:"location"`
	Transaction string `json:"transaction"`

	Amount      int64   `json:"amount"`
	Currency    string  `json:"currency"`
	Description *string `json:"description"`
	Capture     bool    `json:"capture"`
	Authorized  bool    `json:"authorized"`
	Captured    bool    `json:"captured"`
	Card        *Card   `json:"card"`

	Refunded       int64       `json:"refunded"`
	Refunds        *RefundList `json:"refunds"` // TODO: List of refunds object.
	FailureCode    *string     `json:"failure_code"`
	FailureMessage *string     `json:"failure_message"`

	CustomerID   string    `json:"customer"`
	IP           *string   `json:"ip"`
	Dispute      *Dispute  `json:"dispute"`
	Created      time.Time `json:"created"`
	ReturnURI    string    `json:"return_uri"`
	AuthorizeURI string    `json:"authorize_uri"`
}

type ChargeList struct {
	*List
	Data []*Charge `json:"data"`
}
