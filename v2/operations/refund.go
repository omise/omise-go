package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Refund{}, &SearchRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchRefund struct {
}

func (req *SearchRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/refunds/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &ListAllRefunds{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListAllRefunds struct {
}

func (req *ListAllRefunds) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/refunds",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &ListRefunds{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListRefunds struct {
	ChargeID string `json:"-"`
}

func (req *ListRefunds) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &CreateRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateRefund struct {
	ChargeID string `json:"-"`
	Amount int `json:"amount"`
	Metadata map[string]interface{} `json:"metadata"`
	Void bool `json:"void"`
}

func (req *CreateRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &RetrieveRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveRefund struct {
	ChargeID string `json:"-"`
	RefundID string `json:"-"`
}

func (req *RetrieveRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/refunds/" + req.RefundID,
		ContentType: "application/json",
	}
}

