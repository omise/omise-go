package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	receipts, list := &omise.ReceiptList{}, &ListReceipts{
//		List{Offset: 100, Limit: 20},
//	}
//	if e := client.Do(receipts, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("receipts #100-#120:", receipts.Data)
//
type ListReceipts struct {
	List
}

func (req *ListReceipts) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/receipts",
	}
}

// Example:
//
//	rcpt, receipt := &omise.Receipt{}, &RetrieveReceipt{"rcpt_123"}
//	if e := client.Do(rcpt, receipt); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("receipt #123: %#v\n", rcpt)
//
type RetrieveReceipt struct {
	ReceiptID string `query:"-"`
}

func (req *RetrieveReceipt) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/receipts/" + req.ReceiptID,
	}
}
