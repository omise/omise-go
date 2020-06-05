package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

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
	EventID string `json:"-"`
}

func (req *RetrieveEvent) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
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
}

func (req *ListEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/events",
		ContentType: "application/json",
	}
}

