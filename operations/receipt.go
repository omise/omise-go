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

func (req *ListReceipts) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/receipts",
		ContentType: "application/json",
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
	ReceiptID string `json:"-"`
}

func (req *RetrieveReceipt) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/receipts/" + req.ReceiptID,
		ContentType: "application/json",
	}
}
