package omise

// ChargeSchedule represents Charge Schedule object.
// See https://www.omise.co/charge-schedules-api for more information.
type ChargeSchedule struct {
	Base
	Amount int64 `json:"amount"`
	Card string `json:"card"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	DefaultCard bool `json:"default_card"`
	Description *string `json:"description"`
	Metadata map[string]interface{} `json:"metadata"`
}

// ChargeScheduleService represents resource service.
type ChargeScheduleService struct {
	*Client
}

// ChargeSchedule defines resource service.
func (c *Client) ChargeSchedule() *ChargeScheduleService {
	return &ChargeScheduleService{c}
}

