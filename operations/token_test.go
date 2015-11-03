package operations_test

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var CreateTokenOp = &CreateToken{
	Name:            "Chakrit Wichian",
	Number:          "4242424242424242",
	ExpirationMonth: 2,
	ExpirationYear:  time.Now().AddDate(1, 0, 0).Year(),

	SecurityCode: "123",
	City:         "Bangkok",
	PostalCode:   "10240",
}

func TestToken(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	t1, t2 := &omise.Token{}, &omise.Token{}
	if e := client.Do(t1, CreateTokenOp); !a.NoError(t, e) {
		return
	}
	if e := client.Do(t2, &RetreiveToken{ID: t1.ID}); !a.NoError(t, e) {
		return
	}

	a.Equal(t, *t1, *t2)
}
