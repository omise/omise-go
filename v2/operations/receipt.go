package operations


// Example:
//
//	charge, update := &omise.Receipt{}, &ListReceipts{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListReceipts struct {
	Base
}

func (req *ListReceipts) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/receipts",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Receipt{}, &RetrieveReceipt{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveReceipt struct {
	Base
	ReceiptID string `json:"-"`
}

func (req *RetrieveReceipt) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/receipts/" + req.ReceiptID,
		ContentType: "application/json",
	}
}

