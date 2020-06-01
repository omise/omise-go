package omise

// Barcode represents resource object.
type Barcode struct {
	Base
	Image *Document `json:"image"`
	Type string `json:"type"`
}

