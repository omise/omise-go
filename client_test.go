package omise_test

import (
	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	"github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
	"testing"
)

var createTokenOp = &operations.CreateToken{
	Name:            "Chakrit Wichian",
	Number:          "4242424242424242",
	ExpirationMonth: 2,
	ExpirationYear:  2018,
	SecurityCode:    "123",
	City:            "Bangkok",
	PostalCode:      "10210",
}

func TestNewClient(t *testing.T) {
	pkey, skey := testutil.Keys()

	if _, e := NewClient(pkey, skey); !a.NoError(t, e) {
		return
	}
	if _, e := NewClient("", skey); !a.NoError(t, e) {
		return
	}
	if _, e := NewClient(pkey, ""); !a.NoError(t, e) {
		return
	}

	if _, e := NewClient("", ""); a.Error(t, e) {
		a.Equal(t, ErrInvalidKey, e)
	}
	if _, e := NewClient("123", "123"); a.Error(t, e) {
		a.Equal(t, ErrInvalidKey, e)
	}
}
