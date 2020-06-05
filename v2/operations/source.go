package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Source{}, &CreateSource{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateSource struct {
	Amount int `json:"amount"`
	Barcode string `json:"barcode"`
	Currency string `json:"currency"`
	Email string `json:"email"`
	InstallmentTerm int `json:"installment_term"`
	MobileNumber *string `json:"mobile_number"`
	Name string `json:"name"`
	StoreID string `json:"store_id"`
	StoreName string `json:"store_name"`
	TerminalID string `json:"terminal_id"`
	Type *SourceType `json:"type"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

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
//	charge, update := &omise.Source{}, &RetrieveSource{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveSource struct {
	SourceID string `json:"-"`
}

func (req *RetrieveSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/sources/" + req.SourceID,
		ContentType: "application/json",
	}
}

