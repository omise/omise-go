package operations

import (
	"time"
)

// List contains fields that represent parameters common to most list operations. List
// struct is not an operation in and of itself and cannot be used with client.Do directly.
// Use one of the predefined XXList operations instead and supply List struct as the first
// field.
//
// See one of the list operations below for an example.
//
// See the Pagination and Lists documentation at https://www.omise.co/api-pagination for
// more information.
type List struct {
	Offset int
	Limit  int
	From   time.Time
	To     time.Time
	Order  string
}
