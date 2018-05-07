package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	refunds, list := &omise.RefundList{}, &ListRefunds{"chrg_333"}
//	if e := client.Do(refunds, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("refunded on chrg_333:", refunds.Data[0].Amount)
//
type ListRefunds struct {
	ChargeID string
	List
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
//	refund, create := &omise.Refund{}, &CreateRefund{
//		ChargeID: charge.ID,
//		Amount:   charge.Amount >> 1, // half
//	}
//	if e := client.Do(refund, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("refunded half of charge with:", refund.ID)
//
type CreateRefund struct {
	ChargeID string                 `json:"-"`
	Amount   int64                  `json:"amount"`
	Void     bool                   `json:"void,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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
//	refund, retrieve := &omise.Refund{}, &RetrieveRefund{
//		ChargeID: "chrg_123",
//		RefundID: "rfnd_777",
//	}
//	if e := client.Do(refund, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund #777: %#v\n", refund)
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
