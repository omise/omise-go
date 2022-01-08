package omise

// Source represents Omise's source object.
// See https://www.omise.co/source-api for more information.
type Source struct {
	Base
	Type                     string         `json:"type"`
	Flow                     string         `json:"flow"`
	Amount                   int64          `json:"amount"`
	Currency                 string         `json:"currency"`
	ScannableCode            *ScannableCode `json:"scannable_code"`
	References               *References    `json:"references"`
	ZeroInterestInstallments bool           `json:"zero_interest_installments"`
	PlatformType             string         `json:"platform_type"`
}
