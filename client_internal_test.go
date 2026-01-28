package omise

import (
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
