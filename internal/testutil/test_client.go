package testutil

import (
	"os"
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
	r "github.com/stretchr/testify/require"
)

type TestClient struct {
	*testing.T
	*omise.Client
}

func Keys() (string, string) {
	return os.Getenv("OMISE_PKEY"),
		os.Getenv("OMISE_SKEY")
}

func GetCKey() string {
	return os.Getenv("OMISE_CKEY")
}

func NewTestClient(t *testing.T) *TestClient {
	return newTestClient(t, false, false)
}

func NewFixedClient(t *testing.T) *TestClient {
	return newTestClient(t, false, true)
}

func newTestClient(t *testing.T, record, fixed bool) *TestClient {
	client, err := omise.NewClient(Keys())
	r.NoError(t, err)

	switch {
	case fixed:
		fixtures, err := NewFixturesTransport()
		r.NoError(t, err)
		client.Transport = fixtures // override std TLS transport
	case record:
		// WARN: Handle with care!
		recorder, err := NewRecorderTransport()
		r.NoError(t, err)
		client.Transport = recorder
	}

	return &TestClient{t, client}
}

func (tc *TestClient) MustDo(result interface{}, op internal.Operation) {
	r.NoError(tc, tc.Client.Do(result, op))
}
