package omise

// Refund represents Refund object.
// See https://www.omise.co/refunds-api for more information.
type Refund struct {
	Base
	Amount int `json:"amount"`
	Charge string `json:"charge"`
	Currency string `json:"currency"`
	FundingAmount int `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Status *RefundStatus `json:"status"`
	Transaction string `json:"transaction"`
	Voided bool `json:"voided"`
}