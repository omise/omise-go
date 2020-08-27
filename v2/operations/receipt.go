package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	receipt, list := &omise.Receipt{}, &ListReceipts{
//	}
//	if e := client.Do(receipt, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("receipt: %#v\n", receipt)
//
type ListReceipts struct {
	List
}

// Describe
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
//	receipt, retrieve := &omise.Receipt{}, &RetrieveReceipt{
//		ReceiptID: "rcpt_3gqd8bd4h1"
//	}
//	if e := client.Do(receipt, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("receipt: %#v\n", receipt)
//
type RetrieveReceipt struct {
	ReceiptID string `json:"-"`
}

// Describe
func (req *RetrieveReceipt) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/receipts/" + req.ReceiptID,
		ContentType: "application/json",
	}
}

