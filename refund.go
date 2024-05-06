package omise

// Refund represents Omise's refund object.
// See https://www.omise.co/refunds-api for more information.
type Refund struct {
	Base
	Status            string                 `json:"status"`
	Voided            bool                   `json:"voided"`
	FundingAmount     int64                  `json:"funding_amount"`
	Amount            int64                  `json:"amount"`
	Currency          string                 `json:"currency"`
	FundingCurrency   string                 `json:"funding_currency"`
	Charge            string                 `json:"charge"`
	Transaction       string                 `json:"transaction"`
	Metadata          map[string]interface{} `json:"metadata"`
}
