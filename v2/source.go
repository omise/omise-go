package omise

// Source represents Source object.
// See https://www.omise.co/sources-api for more information.
type Source struct {
	Base
	Amount int `json:"amount"`
	Barcode string `json:"barcode"`
	Currency string `json:"currency"`
	Email string `json:"email"`
	Flow *FlowType `json:"flow"`
	InstallmentTerm int `json:"installment_term"`
	Location string `json:"location"`
	MobileNumber string `json:"mobile_number"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	References *PaymentReference `json:"references"`
	ScannableCode *Barcode `json:"scannable_code"`
	StoreID string `json:"store_id"`
	StoreName string `json:"store_name"`
	TerminalID string `json:"terminal_id"`
	Type *SourceType `json:"type"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}
