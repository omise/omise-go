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

// ListService represents resource service.
type ListService struct {
	*Client
}

// ListParams contains fields that represent parameters common to most list operations.
//
// See the Pagination and Lists documentation at https://www.omise.co/api-pagination for
// more information.
type ListParams struct {
	Offset int            `json:"offset,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	From   time.Time      `json:"-"`
	To     time.Time      `json:"-"`
	Order  Ordering `json:"order,omitempty"`
}

// List defines resource service.
func (c *Client) List() *ListService {
	return &ListService{c}
}

