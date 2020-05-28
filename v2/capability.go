package omise

// Capability represents Capability object.
// See https://www.omise.co/capability-api for more information.
type Capability struct {
	Base
	Banks []string `json:"banks"`
	Country string `json:"country"`
	Location string `json:"location"`
	PaymentMethods []PaymentMethod `json:"payment_methods"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}
