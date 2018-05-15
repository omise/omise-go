package operations

import (
	"encoding/json"
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
	Offset int            `json:"offset,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	From   time.Time      `json:"-"`
	To     time.Time      `json:"-"`
	Order  omise.Ordering `json:"order,omitempty"`

	NullableFrom *time.Time `json:"from,omitempty"`
	NullableTo   *time.Time `json:"to,omitempty"`
}

type optionalFieldList List

// MarshalJSON List type
func (l List) MarshalJSON() ([]byte, error) {
	l.SetOptionalField()
	return json.Marshal(optionalFieldList(l))
}

func (l *List) SetOptionalField() {
	if !l.From.IsZero() {
		l.NullableFrom = &l.From
	}
	if !l.To.IsZero() {
		l.NullableTo = &l.To
	}
}
