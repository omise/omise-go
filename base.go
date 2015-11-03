package omise

import "time"

type Base struct {
	Object   string    `json:"object"`
	ID       string    `json:"id"`
	Live     bool      `json:"livemode"`
	Location *string   `json:"location"`
	Created  time.Time `json:"created"`
}
