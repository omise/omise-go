package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
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
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	t1, t2 := &omise.Token{}, &omise.Token{}
	client.MustDo(t1, CreateTokenOp)
	client.MustDo(t2, &RetrieveToken{ID: t1.ID})

	a.Equal(t, *t1, *t2)
}
