package omise

import (
	"time"
)

// List represents resource object.
type List struct {
	Base
	Data []interface{} `json:"data"`
	From time.Time `json:"from"`
	Limit int `json:"limit"`
	Location string `json:"location"`
	Offset int `json:"offset"`
	Order Order `json:"order"`
	To time.Time `json:"to"`
	Total int `json:"total"`
}

