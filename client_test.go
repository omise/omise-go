package omise_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
	"github.com/omise/omise-go/internal/testutil"
	"github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
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

	_, e := NewClient(pkey, skey)
	r.NoError(t, e)
	_, e = NewClient("", skey)
	r.NoError(t, e)
	_, e = NewClient(pkey, "")
	r.NoError(t, e)

	_, e = NewClient("", "")
	r.Error(t, e)
	r.Equal(t, ErrInvalidKey, e)
	_, e = NewClient("123", "123")
	r.Error(t, e)
	r.Equal(t, ErrInvalidKey, e)
}

func TestClient_Request(t *testing.T) {
	pkey, skey := testutil.Keys()
	client, e := NewClient(pkey, skey)
	r.NoError(t, e)

	// use skey for api.omise.co endpoint
	req, e := client.Request(&operations.RetrieveAccount{})
	r.NoError(t, e)

	user, _, ok := req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, user, skey)

	// use pkey for vault.omise.co endopint
	req, e = client.Request(&operations.CreateToken{})
	r.NoError(t, e)

	user, _, ok = req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, user, pkey)

	// use overridden endopint, if specified
	client.Endpoints[internal.API] = "http://api.omise.dev:3000"
	client.Endpoints[internal.Vault] = "http://vault.omise.dev:4500"

	req, e = client.Request(&operations.RetrieveAccount{})
	r.NoError(t, e)
	r.Equal(t, "http://api.omise.dev:3000/account", req.URL.String())

	req, e = client.Request(&operations.CreateToken{})
	r.NoError(t, e)
	r.Equal(t, "http://vault.omise.dev:4500/tokens", req.URL.String())

	// general request properties
	op := &operations.RetrieveAccount{}

	req, e = client.Request(op)
	r.NoError(t, e)
	r.Contains(t, req.Header.Get("User-Agent"), "OmiseGo/")
	r.Contains(t, req.Header.Get("User-Agent"), "Go/go")
	r.Empty(t, req.Header.Get("Omise-Version"), "Omise-Version header sent when APIVersion is not specified.")

	client.GoVersion = "RANDOMXXXVERSION"
	client.APIVersion = "yadda"
	req, e = client.Request(op)
	r.NoError(t, e)
	r.Contains(t, req.Header.Get("User-Agent"), "Go/RANDOMXXXVERSION")
	r.Equal(t, req.Header.Get("Omise-Version"), "yadda")
}

func TestClient_Error(t *testing.T) {
	client, e := NewClient(testutil.Keys())
	r.NoError(t, e)

	e = client.Do(nil, &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/definitely_never_found",
	})
	r.NotNil(t, e)

	err, ok := e.(*Error)
	r.True(t, ok, "error returned is not *omise.Error.")
	r.Equal(t, err.Code, "not_found")

	e = client.Do(nil, &internal.Description{
		Endpoint: internal.Endpoint("virus_endpoint"),
		Method:   "GET",
		Path:     "/",
	})
	r.NotNil(t, e)
	r.IsType(t, ErrInternal(""), e)
}

func TestClient_TransportError(t *testing.T) {
	client := testutil.NewFixedClient(t)

	e := client.Do(&struct{}{}, &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/malformed",
	})
	r.NotNil(t, e)

	err, ok := e.(*ErrTransport)
	r.True(t, ok, "error returned in not *omise.ErrTransport: ")

	_, ok = err.Err.(*json.SyntaxError)
	r.True(t, ok, "error does not wrap *json.SyntaxError")
	r.Contains(t, string(err.Buffer), "not a valid JSON")
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
