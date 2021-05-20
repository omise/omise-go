package omise

import "time"

// References represents Omise's barcode object.
// See https://www.omise.co/sources-api for more information.
type References struct {
	Barcode   string    `json:"barcode"`
	ExpiresAt time.Time `json:"expires_at"`
}
