package omise

type PaymentMethod struct {
	Base
	CardBrands []string `json:"card_brands"`
	Currencies []string `json:"currencies"`
	InstallmentTerms []int `json:"installment_terms"`
	Name string `json:"name"`
}
