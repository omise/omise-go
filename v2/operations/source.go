package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

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
type CreateSource struct {
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
	Type omise.SourceType `json:"type"`
	ZeroInterestInstallments bool `json:"zero_interest_installments,omitempty"`
}

// Describe
func (req *CreateSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/sources",
		ContentType: "application/json",
	}
}

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
type RetrieveSource struct {
	SourceID string `json:"-"`
}

// Describe
func (req *RetrieveSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/sources/" + req.SourceID,
		ContentType: "application/json",
	}
}

