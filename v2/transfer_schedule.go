package omise

// TransferSchedule represents Transfer object.
// See https://www.omise.co/transfer-schedules-api for more information.
type TransferSchedule struct {
	Base
	Amount *int64 `json:"amount"`
	Currency string `json:"currency"`
	PercentageOfBalance *float64 `json:"percentage_of_balance"`
	Recipient string `json:"recipient"`
}

// TransferScheduleService represents resource service.
type TransferScheduleService struct {
	*Client
}

// TransferSchedule defines resource service.
func (c *Client) TransferSchedule() *TransferScheduleService {
	return &TransferScheduleService{c}
}

