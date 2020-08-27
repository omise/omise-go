package omise

// PaymentMethod represents resource object.
type PaymentMethod struct {
	Base
	CardBrands []string `json:"card_brands"`
	Currencies []string `json:"currencies"`
	InstallmentTerms []int64 `json:"installment_terms"`
	Name string `json:"name"`
}

// PaymentMethodService represents resource service.
type PaymentMethodService struct {
	*Client
}

// PaymentMethod defines resource service.
func (c *Client) PaymentMethod() *PaymentMethodService {
	return &PaymentMethodService{c}
}

