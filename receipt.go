package omise

import "time"

// Receipt represents Omise's receipt object.
// See https://www.omise.co/receipt-api for more information.
type Receipt struct {
	Base
	Number                string    `json:"number"`
	Date                  time.Time `json:"date"`
	CustomerName          string    `json:"customer_name"`
	CustomerAddress       string    `json:"customer_address"`
	CustomerTaxID         string    `json:"customer_tax_id"`
	CustomerEmail         string    `json:"customer_email"`
	CustomerStatementName string    `json:"customer_statement_name"`

	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyTaxID   string `json:"company_tax_id"`

	ChargeFee   int64  `json:"charge_fee"`
	VoidedFee   int64  `json:"voided_fee"`
	TransferFee int64  `json:"transfer_fee"`
	SubTotal    int64  `json:"subtotal"`
	VAT         int64  `json:"vat"`
	WHT         int64  `json:"wht"`
	Total       int64  `json:"total"`
	CreditNote  bool   `json:"credit_note"`
	Currency    string `json:"currency"`
}
