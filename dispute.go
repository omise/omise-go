package omise

// Dispute represents Omise's dispute object.
// See https://www.omise.co/disputes-api for more information.
type Dispute struct {
	Base
	Amount   int64                  `json:"amount"`
	Currency string                 `json:"currency"`
	Status   DisputeStatus          `json:"status"`
	Message  string                 `json:"message"`
	Charge   string                 `json:"charge"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
