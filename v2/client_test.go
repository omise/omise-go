package omise_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	. "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2/internal/testutil"
	"github.com/omise/omise-go/v2/operations"
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

	_, err := NewClient(pkey, skey)
	r.NoError(t, err)
	_, err = NewClient("", skey)
	r.NoError(t, err)
	_, err = NewClient(pkey, "")
	r.NoError(t, err)

	_, err = NewClient("", "")
	r.Error(t, err)
	r.Equal(t, ErrInvalidKey, err)
	_, err = NewClient("123", "123")
	r.Error(t, err)
	r.Equal(t, ErrInvalidKey, err)
}

func TestClient_Request(t *testing.T) {
	pkey, skey := testutil.Keys()
	client, err := NewClient(pkey, skey)
	r.NoError(t, err)

	// use skey for api.omise.co endpoint
	req, err := client.Request(&operations.RetrieveAccount{})
	r.NoError(t, err)

	user, _, ok := req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, user, skey)

	// use pkey for vault.omise.co endopint
	req, err = client.Request(&operations.CreateToken{})
	r.NoError(t, err)

	user, _, ok = req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, user, pkey)

	// use overridden endopint, if specified
	client.Endpoints[internal.API] = "http://api.omise.dev:3000"
	client.Endpoints[internal.Vault] = "http://vault.omise.dev:4500"

	req, err = client.Request(&operations.RetrieveAccount{})
	r.NoError(t, err)
	r.Equal(t, "http://api.omise.dev:3000/account", req.URL.String())

	req, err = client.Request(&operations.CreateToken{})
	r.NoError(t, err)
	r.Equal(t, "http://vault.omise.dev:4500/tokens", req.URL.String())

	// general request properties
	desc := &operations.RetrieveAccount{}

	req, err = client.Request(desc)
	r.NoError(t, err)
	r.Contains(t, req.Header.Get("User-Agent"), "OmiseGo/")
	r.Contains(t, req.Header.Get("User-Agent"), "Go/go")
	r.Empty(t, req.Header.Get("Omise-Version"), "Omise-Version header sent when APIVersion is not specified.")

	client.GoVersion = "RANDOMXXXVERSION"
	client.APIVersion = "yadda"
	req, err = client.Request(desc)
	r.NoError(t, err)
	r.Contains(t, req.Header.Get("User-Agent"), "Go/RANDOMXXXVERSION")
	r.Equal(t, req.Header.Get("Omise-Version"), "yadda")
}

func TestClient_Error(t *testing.T) {
	client, err := NewClient(testutil.Keys())
	r.NoError(t, err)

	err = client.Do(nil, &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/definitely_never_found",
	})
	r.NotNil(t, err)

	apiErr, ok := err.(*Error)
	r.True(t, ok, "error returned is not *omise.Error.")
	r.Equal(t, apiErr.Code, "not_found")

	err = client.Do(nil, &internal.Description{
		Endpoint: internal.Endpoint("virus_endpoint"),
		Method:   "GET",
		Path:     "/",
	})
	r.NotNil(t, err)
	r.IsType(t, ErrInternal(""), err)
}

func TestClient_TransportError(t *testing.T) {
	client := testutil.NewFixedClient(t)

	err := client.Do(&struct{}{}, &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/malformed",
	})
	r.NotNil(t, err)

	apiErr, ok := err.(*ErrTransport)
	r.True(t, ok, "error returned in not *omise.ErrTransport: ")

	_, ok = apiErr.Err.(*json.SyntaxError)
	r.True(t, ok, "error does not wrap *json.SyntaxError")
	r.Contains(t, string(apiErr.Buffer), "not a valid JSON")
}

func ExampleClient_Do() {
	// gets your API keys
	pkey, skey := "pkey_test_4yq6tct0llin5nyyi5l", "skey_test_4yq6tct0lblmed2yp5t"

	// creates a client
	client, err := NewClient(pkey, skey)
	if err != nil {
		log.Fatal(err)
	}

	// creates a charge
	charge, create := &Charge{}, &operations.CreateCharge{
		Amount:   100000, // ¥10,000
		Currency: "jpy",
		Card:     "tok_1234",
	}

	// checks for error
	if err := client.Do(charge, create); err != nil {
		if omiseErr, ok := err.(*Error); ok {
			log.Fatal(omiseErr.Code + " " + omiseErr.Message)
		} else {
			log.Fatal("transport error: " + err.Error())
		}
	}

	// verify
	fmt.Printf("authorized charge: %#v\n", charge)
}
