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

func (req *ListEvents) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/events", nil}
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
	EventID string `query:"-"`
}

func (req *RetrieveEvent) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/events/" + req.EventID, nil}
}
