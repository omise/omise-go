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
	Ip                       string         `json:"ip"`
}

type BillingShipping struct {
	Country    string `json:"country"`
	City       string `json:"City"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Street1    string `json:"street1"`
	Street2    string `json:"street2,omitempty"`
}

type Items struct {
	Amount   int64  `json:"amount"`
	Sku      string `json:"sku,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity int16  `json:"quantity,omitempty"`
	Category string `json:"category,omitempty"`
	Brand    string `json:"brand,omitempty"`
	ItemUri  string `json:"item_uri,omitempty"`
	ImageUri string `json:"image_uri,omitempty"`
}
