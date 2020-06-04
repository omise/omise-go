package omise

// ChargeScheduling represents Charge Schedule object.
// See https://www.omise.co/charge-schedules-api for more information.
type ChargeScheduling struct {
	Base
	Amount int `json:"amount"`
	Card string `json:"card"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	DefaultCard bool `json:"default_card"`
	Description *string `json:"description"`
	Metadata map[string]interface{} `json:"metadata"`
}