package schedule

// ChargeDetail represents charge detail for schedule object.
type ChargeDetail struct {
	Amount      int     `json:"amount"`
	Currency    string  `json:"currency"`
	Customer    string  `json:"customer"`
	Card        *string `json:"card"`
	Description string  `json:"description"`
}
