package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Source represents Source object.
// See https://www.omise.co/sources-api for more information.
type Source struct {
	Base
	Amount int64 `json:"amount"`
	Barcode string `json:"barcode"`
	ChargeStatus ChargeStatus `json:"charge_status"`
	Currency string `json:"currency"`
	Email string `json:"email"`
	Flow FlowType `json:"flow"`
	InstallmentTerm int64 `json:"installment_term"`
	Location string `json:"location"`
	MobileNumber string `json:"mobile_number"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	References *PaymentReference `json:"references"`
	ScannableCode *Barcode `json:"scannable_code"`
	StoreID string `json:"store_id"`
	StoreName string `json:"store_name"`
	TerminalID string `json:"terminal_id"`
	Type SourceType `json:"type"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

// SourceService represents resource service.
type SourceService struct {
	*Client
}

// Source defines resource service.
func (c *Client) Source() *SourceService {
	return &SourceService{c}
}

// Create creates source
//
// Example:
//
//	source, create := &omise.Source{}, &CreateSource{
//		Amount: 10000
//		Currency: "thb"
//		Type: nil
//	}
//	if e := client.Do(source, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("source: %#v\n", source)
//
func (s *SourceService) Create(ctx context.Context, params *CreateSourceParams) (*Source, error) {
	result := &Source{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateSourceParams params object.
type CreateSourceParams struct {
	Amount int64 `json:"amount"`
	Barcode string `json:"barcode,omitempty"`
	Currency string `json:"currency"`
	Email string `json:"email,omitempty"`
	InstallmentTerm int64 `json:"installment_term,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	Name string `json:"name,omitempty"`
	StoreID string `json:"store_id,omitempty"`
	StoreName string `json:"store_name,omitempty"`
	TerminalID string `json:"terminal_id,omitempty"`
	Type SourceType `json:"type"`
	ZeroInterestInstallments bool `json:"zero_interest_installments,omitempty"`
}

// Describe describes structure of request
func (req *CreateSourceParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/sources",
		ContentType: "application/json",
	}
}

// Retrieve retrieves source
//
// Example:
//
//	source, retrieve := &omise.Source{}, &RetrieveSource{
//		SourceID: "sourc_4hrdn6jt3g"
//	}
//	if e := client.Do(source, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("source: %#v\n", source)
//
func (s *SourceService) Retrieve(ctx context.Context, params *RetrieveSourceParams) (*Source, error) {
	result := &Source{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveSourceParams params object.
type RetrieveSourceParams struct {
	SourceID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveSourceParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/sources/" + req.SourceID,
		ContentType: "application/json",
	}
}

