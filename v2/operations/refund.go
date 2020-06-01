package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	refund, listRefunds := &omise.Refund{}, &ListAllRefunds{
//	}
//	if e := client.Do(refund, listRefunds); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
type ListAllRefunds struct {
	List
}

// Describe
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
//	refund, list := &omise.Refund{}, &ListRefunds{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(refund, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
type ListRefunds struct {
	ChargeID string `json:"-"`
	List
}

// Describe
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
//	refund, create := &omise.Refund{}, &CreateRefund{
//		ChargeID: "chrg_8b3g63gq2f"
//		Amount: 10000
//	}
//	if e := client.Do(refund, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
type CreateRefund struct {
	ChargeID string `json:"-"`
	Amount int64 `json:"amount"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Void bool `json:"void,omitempty"`
}

// Describe
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
//	refund, retrieve := &omise.Refund{}, &RetrieveRefund{
//		ChargeID: "chrg_8b3g63gq2f"
//		RefundID: "refun_3gqd1e6jt9"
//	}
//	if e := client.Do(refund, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
type RetrieveRefund struct {
	ChargeID string `json:"-"`
	RefundID string `json:"-"`
}

// Describe
func (req *RetrieveRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/refunds/" + req.RefundID,
		ContentType: "application/json",
	}
}

