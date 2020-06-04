package operations


// Example:
//
//	charge, update := &omise.Transaction{}, &RetrieveTransaction{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveTransaction struct {
	Base
	TransactionID string `json:"-"`
}

func (req *RetrieveTransaction) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transactions/" + req.TransactionID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transaction{}, &ListTransactions{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListTransactions struct {
	Base
}

func (req *ListTransactions) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transactions",
		ContentType: "application/json",
	}
}

