package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	events, list := &omise.EventList{}, &ListEvents{}
//	if e := client.Do(events, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of events:", len(events.Data))
//
type ListEvents struct {
	List
}

func (req *ListEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/events",
		ContentType: "application/json",
	}
}

// Example:
//
//	event, retrieve := &omise.Event{}, &RetrieveEvent{"evnt_123"}
//	if e := client.Do(event, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("evnt_123: %#v\n", event)
//
type RetrieveEvent struct {
	EventID string `json:"-"`
}

func (req *RetrieveEvent) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/events/" + req.EventID,
	}
}
