package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Event represents Event object.
// See https://www.omise.co/events-api for more information.
type Event struct {
	Base
	Data interface{} `json:"data"`
	Key string `json:"key"`
	Location string `json:"location"`
	WebhookDeliveries []interface{} `json:"webhook_deliveries"`
}

// EventService represents resource service.
type EventService struct {
	*Client
}

// Event defines resource service.
func (c *Client) Event() *EventService {
	return &EventService{c}
}

// Retrieve retrieves event
//
// Example:
//
//	event, err := client.Event().Retrieve(ctx, &RetrieveEventParams{
//		EventID: "evnt_d7akud9cm5"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
func (s *EventService) Retrieve(ctx context.Context, params *RetrieveEventParams) (*Event, error) {
	result := &Event{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveEventParams params object.
type RetrieveEventParams struct {
	EventID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveEventParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/events/" + req.EventID,
		ContentType: "application/json",
	}
}

// List lists events
//
// Example:
//
//	event, err := client.Event().List(ctx, &ListEventsParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
func (s *EventService) List(ctx context.Context, params *ListEventsParams) (*EventList, error) {
	result := &EventList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListEventsParams params object.
type ListEventsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListEventsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/events",
		ContentType: "application/json",
	}
}

