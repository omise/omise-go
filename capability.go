package omise

// Capability represents Omise's capability object.
// See https://www.omise.co/capability-api for more information.
type Capability struct {
	Object                   string          `json:"object"`
	Location                 string          `json:"location"`
	Banks                    []string        `json:"banks"`
	Country                  string          `json:"country"`
	PaymentMethods           []PaymentMethod `json:"payment_methods"`
	ZeroInterestInstallments bool            `json:"zero_interest_installments"`
}

// PaymentMethod represents a payment method in the Omise's capability object.
type PaymentMethod struct {
	Object           string   `json:"object"`
	Name             string   `json:"name"`
	Currencies       []string `json:"currencies"`
	CardBrands       []string `json:"card_brands"`
	InstallmentTerms []int    `json:"installment_terms"`
	Banks            []Bank   `json:"banks"`
}

// Bank represents a bank object in the payment method of the capability object.
type Bank struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
