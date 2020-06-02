package omise

// Forex represents Forex object.
// See https://www.omise.co/forex-api for more information.
type Forex struct {
	BaseModel
	Base string `json:"base"`
	Location string `json:"location"`
	Quote string `json:"quote"`
	Rate int `json:"rate"`
}
