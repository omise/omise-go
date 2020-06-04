package operations


// Example:
//
//	charge, update := &omise.Chain{}, &ListChains{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListChains struct {
	Base
}

func (req *ListChains) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/chains",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Chain{}, &RetrieveChain{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveChain struct {
	Base
	ChainID string `json:"-"`
}

func (req *RetrieveChain) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/chains/" + req.ChainID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Chain{}, &RevokeChain{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RevokeChain struct {
	Base
	ChainID string `json:"-"`
}

func (req *RevokeChain) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/chains/" + req.ChainID + "/revoke",
		ContentType: "application/json",
	}
}

