package omise

// Token represents Token object.
// See https://www.omise.co/tokens-api for more information.
type Token struct {
	Base
	Card *Card `json:"card"`
	Location string `json:"location"`
	Used bool `json:"used"`
}

