package omise

// Capability represents Omise's capability object.
// See https://www.omise.co/capability-api for more information.
type Capability struct {
	Base
	Banks                    []string            `json:"banks"`
	PaymentMethods           []PaymentMethodType `json:"payment_methods"`
	Country                  string              `json:"country"`
	ZeroInterestInstallments bool                `json:"zero_interest_installments"`
}

// PaymentMethodType represents payment method struct
type PaymentMethodType struct {
	Base
	Name             string   `json:"name"`
	Currencies       []string `json:"currencies"`
	CardBrands       []string `json:"card_brands"`
	InstallmentTerms *int     `json:"installment_terms"`
}
