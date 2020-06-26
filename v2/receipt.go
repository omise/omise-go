package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

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

// ReceiptService represents resource service.
type ReceiptService struct {
	*Client
}

// Receipt defines resource service.
func (c *Client) Receipt() *ReceiptService {
	return &ReceiptService{c}
}

// List lists receipts
//
// Example:
//
//	receipt, list := &omise.ReceiptList{}, &ListReceipts{
//	}
//	if e := client.Do(receipt, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("receipt: %#v\n", receipt)
//
func (s *ReceiptService) List(ctx context.Context, params *ListReceiptsParams) (*ReceiptList, error) {
	result := &ReceiptList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListReceiptsParams params object.
type ListReceiptsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListReceiptsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/receipts",
		ContentType: "application/json",
	}
}

// Retrieve retrieves receipt
//
// Example:
//
//	receipt, retrieve := &omise.Receipt{}, &RetrieveReceipt{
//		ReceiptID: "rcpt_3gqd8bd4h1"
//	}
//	if e := client.Do(receipt, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("receipt: %#v\n", receipt)
//
func (s *ReceiptService) Retrieve(ctx context.Context, params *RetrieveReceiptParams) (*Receipt, error) {
	result := &Receipt{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveReceiptParams params object.
type RetrieveReceiptParams struct {
	ReceiptID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveReceiptParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/receipts/" + req.ReceiptID,
		ContentType: "application/json",
	}
}

