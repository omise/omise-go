package omise

import (
	"time"
)

// Dispute represents Dispute object.
// See https://www.omise.co/disputes-api for more information.
type Dispute struct {
	Base
	AdminMessage *string `json:"admin_message"`
	Amount int64 `json:"amount"`
	Charge string `json:"charge"`
	ClosedAt *time.Time `json:"closed_at"`
	Currency string `json:"currency"`
	Documents *DocumentList `json:"documents"`
	FundingAmount int64 `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Location string `json:"location"`
	Message string `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
	ReasonCode DisputeReasonCode `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
	Status DisputeStatus `json:"status"`
	Transactions []Transaction `json:"transactions"`
}

