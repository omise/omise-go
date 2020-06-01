package omise

// PaymentMethod represents resource object.
type PaymentMethod struct {
	Base
	CardBrands []string `json:"card_brands"`
	Currencies []string `json:"currencies"`
	InstallmentTerms []int64 `json:"installment_terms"`
	Name string `json:"name"`
}

