package omise

import "time"

// Charge represents Omise's charge object.
// See https://www.omise.co/charges-api for more information.
type Charge struct {
	Base
	Status            ChargeStatus      `json:"status"`
	Amount            int64             `json:"amount"`
	AuthorizationType AuthorizationType `json:"authorization_type"`
	AuthorizedAmount  int64             `json:"authorized_amount"`
	CapturedAmount    int64             `json:"captured_amount"`
	Currency          string            `json:"currency"`
	Description       *string           `json:"description"`

	Capture    bool `json:"capture"`
	Authorized bool `json:"authorized"`
	Reversed   bool `json:"reversed"`
	Paid       bool `json:"paid"`

	Transaction string `json:"transaction"`
	Card        *Card  `json:"card"`

	RefundedAmount int64       `json:"refunded_amount"`
	Refunds        *RefundList `json:"refunds"`
	FailureCode    *string     `json:"failure_code"`
	FailureMessage *string     `json:"failure_message"`

	CustomerID string   `json:"customer"`
	IP         *string  `json:"ip"`
	Dispute    *Dispute `json:"dispute"`

	ReturnURI    string `json:"return_uri"`
	AuthorizeURI string `json:"authorize_uri"`

	SourceOfFund SourceOfFunds `json:"source_of_fund"`

	Source    *Source                `json:"source"`
	Metadata  map[string]interface{} `json:"metadata"`
	ExpiresAt time.Time              `json:"expires_at"`

	MerchantAdvice     string              `json:"merchant_advice"`
	MerchantAdviceCode string              `json:"merchant_advice_code"`
	Missing3DSFields   []string            `json:"missing_3ds_fields"`
	AuthenticatedBy    *AuthenticationType `json:"authenticated_by"`
}

type TransactionIndicator string

const (
	MIT TransactionIndicator = "MIT"
	CIT TransactionIndicator = "CIT"
)

type RecurringReason string

const (
	Blank           RecurringReason = ""
	Unscheduled     RecurringReason = "unscheduled"
	StandingOrder   RecurringReason = "standing_order"
	Subscription    RecurringReason = "subscription"
	Installment     RecurringReason = "installment"
	PartialShipment RecurringReason = "partial_shipment"
	DelayedCharge   RecurringReason = "delayed_charge"
	NoShow          RecurringReason = "no_show"
	Resubmission    RecurringReason = "resubmission"
)

type PlatformFee struct {
	Fixed      int32   `json:"fixed,omitempty"`
	Percentage float32 `json:"percentage,omitempty"`
	Amount     int64   `json:"Amount,omitempty"`
}
