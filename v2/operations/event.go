package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	event, retrieve := &omise.Event{}, &RetrieveEvent{
//		EventID: "evnt_d7akud9cm5"
//	}
//	if e := client.Do(event, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
type RetrieveEvent struct {
	EventID string `json:"-"`
}

// Describe
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
//	event, list := &omise.Event{}, &ListEvents{
//	}
//	if e := client.Do(event, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
type ListEvents struct {
	List
}

// Describe
func (req *ListEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/events",
		ContentType: "application/json",
	}
}

