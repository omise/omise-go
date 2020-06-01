package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	omise "github.com/omise/omise-go/v2"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestListMarshal(t *testing.T) {
	testdata := []struct {
		req      List
		expected string
	}{
		{
			req:      List{},
			expected: `{}`,
		},
		{
			req: List{
				Offset: 1,
			},
			expected: `{"offset":1}`,
		},
		{
			req: List{
				Limit: 5,
			},
			expected: `{"limit":5}`,
		},
		{
			req: List{
				From: time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: `{"from":"2017-05-01T00:00:00Z"}`,
		},
		{
			req: List{
				To: time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: `{"to":"2017-05-01T00:00:00Z"}`,
		},
		{
			req: List{
				Order: omise.Chronological,
			},
			expected: `{"order":"chronological"}`,
		},
		{
			req: List{
				Offset: 1,
				Limit:  5,
				From:   time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC),
				To:     time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
				Order:  omise.Chronological,
			},
			expected: `{"offset":1,"limit":5,"order":"chronological","from":"2017-05-01T00:00:00Z","to":"2017-10-01T00:00:00Z"}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}
