package omise

// PlatformFee represents resource object.
type PlatformFee struct {
	Base
	Amount int64 `json:"amount"`
	Fixed int64 `json:"fixed"`
	Percentage float64 `json:"percentage"`
}

