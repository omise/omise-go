package operations


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
	Base
}

func (req *SearchRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/refunds/search",
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
	Base
}

func (req *ListRefunds) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/refunds",
		ContentType: "application/json",
	}
}

