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
	ChargeID string `query:"-"`
	List
}

func (req *ListRefunds) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID + "/refunds", nil}
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
	ChargeID string `query:"-"`
	Amount   int64
}

func (req *CreateRefund) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/charges/" + req.ChargeID + "/refunds", nil}
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
	ChargeID string `query:"-"`
	RefundID string `query:"-"`
}

func (req *RetrieveRefund) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID + "/refunds/" + req.RefundID, nil}
}
