package omise

type PlatformFee struct {
	Base
	Amount int `json:"amount"`
	Fixed int `json:"fixed"`
	Percentage int `json:"percentage"`
}