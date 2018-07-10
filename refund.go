package omise

// Refund represents Omise's refund object.
// See https://www.omise.co/refunds-api for more information.
type Refund struct {
	Base
	Amount      int64                  `json:"amount"`
	Currency    string                 `json:"currency"`
	Charge      string                 `json:"charge"`
	Transaction string                 `json:"transaction"`
	Metadata    map[string]interface{} `json:"metadata"`
}
