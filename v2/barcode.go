package omise

type Barcode struct {
	Base
	Image *Document `json:"image"`
	Type string `json:"type"`
}