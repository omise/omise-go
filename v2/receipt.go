package omise

// Receipt represents Receipt object.
// See https://www.omise.co/receipts-api for more information.
type Receipt struct {
	Base
	AdjustmentTransaction *Transaction `json:"adjustment_transaction"`
	ChargeFee int64 `json:"charge_fee"`
	CompanyAddress string `json:"company_address"`
	CompanyName string `json:"company_name"`
	CompanyTaxID string `json:"company_tax_id"`
	CreditNote bool `json:"credit_note"`
	Currency string `json:"currency"`
	CustomerAddress string `json:"customer_address"`
	CustomerEmail string `json:"customer_email"`
	CustomerName string `json:"customer_name"`
	CustomerStatementName string `json:"customer_statement_name"`
	CustomerTaxID string `json:"customer_tax_id"`
	IssuedOn Date `json:"issued_on"`
	Location string `json:"location"`
	Number string `json:"number"`
	Subtotal int64 `json:"subtotal"`
	Total int64 `json:"total"`
	TransferFee int64 `json:"transfer_fee"`
	Vat int64 `json:"vat"`
	VoidedFee int64 `json:"voided_fee"`
	Wht int64 `json:"wht"`
}

