package omise

// ChargeScheduleDetail represents charge detail for schedule object.
type ChargeScheduleDetail struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	Card     string `json:"card"`
}
