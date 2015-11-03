package operations

import (
	"time"
)

// REF: https://www.omise.co/api-pagination
type ListParams struct {
	Offset int
	Limit  int
	From   time.Time
	To     time.Time
	Order  string
}
