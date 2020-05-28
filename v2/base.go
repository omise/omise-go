package omise

import "time"

// Base structure contains fields that are common to objects returned by the Omise's REST
// API.
type Base struct {
	Object   string    `json:"object"`
	ID       string    `json:"id"`
	Live     bool      `json:"livemode"`
	Location *string   `json:"location"`
	Created  time.Time `json:"created"`
}

// Deletion struct is used to receive deletion responses from delete operations.
type Deletion struct {
	Base
	Deleted bool `json:"deleted"`
}
