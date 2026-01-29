package omise

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/omise/omise-go/internal"
	r "github.com/stretchr/testify/require"
)

type captureTransport struct {
	req *http.Request
}

func (t *captureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req

	return &http.Response{
		StatusCode:    http.StatusOK,
		Body:          io.NopCloser(strings.NewReader(`{}`)),
		Header:        make(http.Header),
		ContentLength: 2,
		Request:       req,
	}, nil
}

func TestDoWithKeysUsesOverrideSecret(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	capture := &captureTransport{}
	client.Transport = capture

	overridePKey := "pkey_test_override"
	overrideSKey := "skey_test_override"

	err = client.DoWithKeys(&struct{}{}, overridePKey, overrideSKey, &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/dummy",
	})
	r.NoError(t, err)

	user, _, ok := capture.req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, overrideSKey, user) // secret key should be used for API endpoint
}

func TestDoWithKeysUsesOverridePublic(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	capture := &captureTransport{}
	client.Transport = capture

	overridePKey := "pkey_test_override"
	overrideSKey := "skey_test_override"

	err = client.DoWithKeys(&struct{}{}, overridePKey, overrideSKey, &internal.Description{
		Endpoint: internal.Vault,
		Method:   "GET",
		Path:     "/dummy",
	})
	r.NoError(t, err)

	user, _, ok := capture.req.BasicAuth()
	r.True(t, ok)
	r.Equal(t, overridePKey, user) // public key should be used for Vault endpoint
}

type badMarshalOp struct {
	internal.Description
	Bad func()
}

func (o *badMarshalOp) Describe() *internal.Description { return &o.Description }

type badMethodOp struct {
	internal.Description
}

func (o *badMethodOp) Describe() *internal.Description { return &o.Description }

type badEndpointOp struct {
	internal.Description
}

func (o *badEndpointOp) Describe() *internal.Description { return &o.Description }

type errorTransport struct {
	err error
}

func (t *errorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, t.err
}

func TestNewClientRejectsInvalidSecretPrefix(t *testing.T) {
	_, err := NewClient("pkey_test_good", "secret_without_prefix")
	r.Equal(t, ErrInvalidKey, err)
}

func TestRequestJSONMarshalError(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	_, err = client.Request(&badMarshalOp{
		Description: internal.Description{Endpoint: internal.API, Method: "GET", Path: "/"},
		Bad: func() {
			// intentionally empty: presence of a func field forces json.Marshal to error
		},
	})
	r.Error(t, err)
}

func TestRequestWithKeysInvalidMethod(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	_, err = client.requestWithKeys(&badMethodOp{
		Description: internal.Description{Endpoint: internal.API, Method: "GET\r\n", Path: "/"},
	}, "pkey_test_default", "skey_test_default")
	r.Error(t, err)
}

func TestDoWithKeysUnknownEndpoint(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	err = client.DoWithKeys(&struct{}{}, "pkey_test_default", "skey_test_default", &badEndpointOp{
		Description: internal.Description{Endpoint: internal.Endpoint("bad"), Method: "GET", Path: "/"},
	})
	r.Error(t, err)
}

func TestDoWithRequestTransportError(t *testing.T) {
	client, err := NewClient("pkey_test_default", "skey_test_default")
	r.NoError(t, err)

	client.Transport = &errorTransport{err: fmt.Errorf("boom")}

	err = client.DoWithKeys(&struct{}{}, "pkey_test_default", "skey_test_default", &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/",
	})
	r.Error(t, err)
}
