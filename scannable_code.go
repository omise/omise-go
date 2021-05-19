package omise

// Scannable code represents Omise's barcode object.
// See https://www.omise.co/sources-api for more information.
type ScannableCode struct {
	Object string    `json:"object"`
	Type   string    `json:"type"`
	Image  *Document `json:"image"`
}
