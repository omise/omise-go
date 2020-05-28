package omise

import (
	"time"
)

// Charge represents Charge object.
// See https://www.omise.co/charges-api for more information.
type Charge struct {
	Base
	Amount int `json:"amount"`
	AuthorizeURI *string `json:"authorize_uri"`
	Authorized bool `json:"authorized"`
	Capturable bool `json:"capturable"`
	Capture bool `json:"capture"`
	Card *Card `json:"card"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	Description *string `json:"description"`
	Disputable bool `json:"disputable"`
	Dispute *Dispute `json:"dispute"`
	Expired bool `json:"expired"`
	ExpiredAt *time.Time `json:"expired_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	FailureCode *ChargeFailureCode `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Fee int `json:"fee"`
	FeeVat int `json:"fee_vat"`
	FundingAmount int `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Interest int `json:"interest"`
	InterestVat int `json:"interest_vat"`
	IP *string `json:"ip"`
	Link string `json:"link"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Net int `json:"net"`
	Paid bool `json:"paid"`
	PaidAt time.Time `json:"paid_at"`
	PlatformFee *PlatformFee `json:"platform_fee"`
	Refundable bool `json:"refundable"`
	RefundedAmount int `json:"refunded_amount"`
	Refunds *RefundList `json:"refunds"`
	ReturnURI *string `json:"return_uri"`
	Reversed bool `json:"reversed"`
	ReversedAt *time.Time `json:"reversed_at"`
	Reversible bool `json:"reversible"`
	Schedule string `json:"schedule"`
	Source *Source `json:"source"`
	Status *ChargeStatus `json:"status"`
	Transaction string `json:"transaction"`
	Voided bool `json:"voided"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}
