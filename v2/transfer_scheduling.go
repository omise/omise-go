package omise

// TransferScheduling represents Transfer object.
// See https://www.omise.co/transfer-schedules-api for more information.
type TransferScheduling struct {
	Base
	Amount *int64 `json:"amount"`
	Currency string `json:"currency"`
	PercentageOfBalance *float64 `json:"percentage_of_balance"`
	Recipient string `json:"recipient"`
}

