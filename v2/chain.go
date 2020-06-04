package omise

// Chain represents Chain object.
// See https://www.omise.co/chains-api for more information.
type Chain struct {
	Base
	Email string `json:"email"`
	Key string `json:"key"`
	Location string `json:"location"`
	Revoked bool `json:"revoked"`
}