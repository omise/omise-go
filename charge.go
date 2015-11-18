package omise

import (
	"time"
)

// Charge represents Omise's charge object.
// See https://www.omise.co/charges-api for more information.
type Charge struct {
	Base
	Transaction string  `json:"transaction"`
	Amount      int64   `json:"amount" pretty:""`
	Currency    string  `json:"currency" pretty:""`
	Description *string `json:"description"`
	Capture     bool    `json:"capture" pretty:""`
	Authorized  bool    `json:"authorized" pretty:""`
	Captured    bool    `json:"captured" pretty:""`
	Card        *Card   `json:"card"`

	Refunded       int64       `json:"refunded"`
	Refunds        *RefundList `json:"refunds"`
	FailureCode    *string     `json:"failure_code"`
	FailureMessage *string     `json:"failure_message"`

	CustomerID string   `json:"customer"`
	IP         *string  `json:"ip"`
	Dispute    *Dispute `json:"dispute"`

	Created      time.Time `json:"created"`
	ReturnURI    string    `json:"return_uri"`
	AuthorizeURI string    `json:"authorize_uri"`
}
