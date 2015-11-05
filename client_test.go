package omise_test

import (
	"fmt"
	"log"
	"testing"

	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
	"github.com/omise/omise-go/internal/testutil"
	"github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
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

func TestClient_Error(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	e = client.Do(nil, &internal.Op{
		internal.API, "GET", "/definitely_never_found", nil,
	})
	if a.NotNil(t, e) {
		err, ok := e.(*Error)
		if a.True(t, ok, "error returned is not *omise.Error.") {
			a.Equal(t, err.Code, "not_found")
		}
	}

	e = client.Do(nil, &internal.Op{
		internal.Endpoint("virus_endpoint"), "GET", "/", nil,
	})
	if a.NotNil(t, e) {
		a.IsType(t, ErrInternal(""), e)
	}
}

func ExampleClient_Do() {
	// gets your API keys
	pkey, skey := "pkey_test_4yq6tct0llin5nyyi5l", "skey_test_4yq6tct0lblmed2yp5t"

	// creates a client
	client, e := NewClient(pkey, skey)
	if e != nil {
		log.Fatal(e)
	}

	// creates a charge
	charge, create := &Charge{}, &operations.CreateCharge{
		Amount:   100000, // ¥10,000
		Currency: "jpy",
		Card:     "tok_1234",
	}

	// checks for error
	if e := client.Do(charge, create); e != nil {
		if omiseErr, ok := e.(*Error); ok {
			log.Fatal(omiseErr.Code + " " + omiseErr.Message)
		} else {
			log.Fatal("transport error: " + e.Error())
		}
	}

	// verify
	fmt.Printf("authorized charge: %#v\n", charge)
}
