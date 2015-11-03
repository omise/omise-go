package omise

import "time"

type Token struct {
	Base
	Used    bool      `json:"used"`
	Card    *Card     `json:"card"`
	Created time.Time `json:"created"`
}
