package omise

// Balance represents Balance object.
// See https://www.omise.co/balance-api for more information.
type Balance struct {
	Base
	Currency string `json:"currency"`
	Location string `json:"location"`
	Reserve int `json:"reserve"`
	Total int `json:"total"`
	Transferable int `json:"transferable"`
}