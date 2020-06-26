package omise

// Barcode represents resource object.
type Barcode struct {
	Base
	Image *Document `json:"image"`
	Type string `json:"type"`
}

// BarcodeService represents resource service.
type BarcodeService struct {
	*Client
}

// Barcode defines resource service.
func (c *Client) Barcode() *BarcodeService {
	return &BarcodeService{c}
}

