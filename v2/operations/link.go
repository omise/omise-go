package operations


// Example:
//
//	charge, update := &omise.Link{}, &SearchLink{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchLink struct {
	Base
}

func (req *SearchLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/links/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Link{}, &DestroyLink{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyLink struct {
	Base
	LinkID string `json:"-"`
}

func (req *DestroyLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Link{}, &RetrieveLink{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveLink struct {
	Base
	LinkID string `json:"-"`
}

func (req *RetrieveLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Link{}, &ListLinks{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListLinks struct {
	Base
}

func (req *ListLinks) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/links",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Link{}, &CreateLink{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateLink struct {
	Base
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Multiple bool `json:"multiple"`
	Title string `json:"title"`
}

func (req *CreateLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/links",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &ListLinksCharges{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListLinksCharges struct {
	Base
	LinkID string `json:"-"`
}

func (req *ListLinksCharges) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/links/" + req.LinkID + "/charges",
		ContentType: "application/json",
	}
}

