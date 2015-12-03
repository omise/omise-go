package omise

import (
	"encoding/json"
	"net/http"
)

// EventHandler is an interface for handling events from the webhook http.Handler.
type EventHandler interface {
	HandleEvent(http.ResponseWriter, *http.Request, *Event)
}

// EventHandlerFunc lets you use a plain function as an EventHandler type.
type EventHandlerFunc func(http.ResponseWriter, *http.Request, *Event)

func (f EventHandlerFunc) HandleEvent(resp http.ResponseWriter, req *http.Request, event *Event) {
	f(resp, req, event)
}

type webhookHTTPHandler struct {
	eventHandler EventHandler
}

func (h *webhookHTTPHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	event := &Event{}
	if e := json.NewDecoder(req.Body).Decode(event); e != nil {
		http.Error(resp, e.Error(), http.StatusBadRequest)
		return
	}

	h.eventHandler.HandleEvent(resp, req, event)
}

// WebhookHTTPHandler creates an http.Handler that you can use to receive Omise's webhook
// API calls for events. The returned handler will automatically consume the request body
// and unmarshals an Event object from JSON for you.
//
// See https://www.omise.co/api-webhooks for more information.
func WebhookHTTPHandler(handler EventHandler) http.Handler {
	return &webhookHTTPHandler{handler}
}
