package schedule

// TransferDetail represents transfer detail for schedule object.
type TransferDetail struct {
	Recipient           string   `json:"recipient"`
	Amount              *int     `json:"amount"`
	PercentageOfBalance *float64 `json:"percentage_of_balance"`
	Currency            string   `json:"currency"`
}
