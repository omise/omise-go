package operations


// Example:
//
//	charge, update := &omise.Capability{}, &RetrieveCapability{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveCapability struct {
	Base
}

func (req *RetrieveCapability) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/capability",
		ContentType: "application/json",
	}
}

