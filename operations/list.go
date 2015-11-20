package operations

import (
	"time"

	"github.com/omise/omise-go"
)

// List contains fields that represent parameters common to most list operations. List
// struct is not an operation in and of itself and cannot be used with client.Do directly.
// Use one of the predefined XXXList operations defined blow instead and supply List
// struct as the first field.
//
// See the Pagination and Lists documentation at https://www.omise.co/api-pagination for
// more information.
type List struct {
	Offset int
	Limit  int
	From   time.Time
	To     time.Time
	Order  omise.Ordering
}
