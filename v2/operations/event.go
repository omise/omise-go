package operations


// Example:
//
//	charge, update := &omise.Event{}, &RetrieveEvent{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveEvent struct {
	Base
	EventID string `json:"-"`
}

func (req *RetrieveEvent) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/events/" + req.EventID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Event{}, &ListEvents{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListEvents struct {
	Base
}

func (req *ListEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/events",
		ContentType: "application/json",
	}
}

