package omise

// Balance represents Balance object.
// See https://www.omise.co/balance-api for more information.
type Balance struct {
	Base
	Currency string `json:"currency"`
	Location string `json:"location"`
	Reserve int64 `json:"reserve"`
	Total int64 `json:"total"`
	Transferable int64 `json:"transferable"`
}

