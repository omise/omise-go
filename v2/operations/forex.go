package operations


// Example:
//
//	charge, update := &omise.Forex{}, &RetrieveForex{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveForex struct {
	Base
	Currency string `json:"-"`
}

func (req *RetrieveForex) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/forex/" + req.Currency,
		ContentType: "application/json",
	}
}

