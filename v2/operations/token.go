package operations


// Example:
//
//	charge, update := &omise.Token{}, &CreateToken{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateToken struct {
	Base
	Card *Card `json:"card"`
}

func (req *CreateToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/tokens",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Token{}, &RetrieveToken{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveToken struct {
	Base
	TokenID string `json:"-"`
}

func (req *RetrieveToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/tokens/" + req.TokenID,
		ContentType: "application/json",
	}
}

