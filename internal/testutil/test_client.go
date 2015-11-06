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
	return os.Getenv("OMISE_PUB_KEY"),
		os.Getenv("OMISE_SECRET_KEY")
}

func NewTestClient(t *testing.T) *TestClient {
	return newTestClient(t, false)
}

func NewFixedClient(t *testing.T) *TestClient {
	return newTestClient(t, true)
}

func newTestClient(t *testing.T, fixed bool) *TestClient {
	client, e := omise.NewClient(Keys())
	r.NoError(t, e)

	if fixed {
		fixtures, e := NewFixturesTransport("")
		r.NoError(t, e)
		client.Transport = fixtures // override std TLS transport
	}

	return &TestClient{t, client}
}

func (tc *TestClient) MustDo(result interface{}, op internal.Operation) {
	r.NoError(tc, tc.Client.Do(result, op))
}
