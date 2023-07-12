package omise

import "time"

// Charge represents Omise's charge object.
// See https://www.omise.co/charges-api for more information.
type Charge struct {
	Base
	Status      ChargeStatus `json:"status"`
	Amount      int64        `json:"amount"`
	Currency    string       `json:"currency"`
	Description *string      `json:"description"`

	Capture    bool `json:"capture"`
	Authorized bool `json:"authorized"`
	Reversed   bool `json:"reversed"`
	Paid       bool `json:"paid"`

	Transaction string `json:"transaction"`
	Card        *Card  `json:"card"`

	Refunded       int64       `json:"refunded"`
	Refunds        *RefundList `json:"refunds"`
	FailureCode    *string     `json:"failure_code"`
	FailureMessage *string     `json:"failure_message"`

	CustomerID string   `json:"customer"`
	IP         *string  `json:"ip"`
	Dispute    *Dispute `json:"dispute"`

	ReturnURI    string `json:"return_uri"`
	AuthorizeURI string `json:"authorize_uri"`

	SourceOfFund SourceOfFunds `json:"source_of_fund"`
	Offsite      OffsiteTypes  `json:"offsite"`

	Source    *Source                `json:"source"`
	Metadata  map[string]interface{} `json:"metadata"`
	ExpiresAt time.Time              `json:"expires_at"`
}
