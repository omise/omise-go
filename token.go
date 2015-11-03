package omise

import "time"

type Token struct {
	ID       string `json:"id"`
	Live     bool   `json:"livemode"`
	Location string `json:"location"`

	Used    bool      `json:"used"`
	Card    *Card     `json:"card"`
	Created time.Time `json:"created"`
}
