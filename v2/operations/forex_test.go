package operations_test

import (
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/assert"
)

func TestForex(t *testing.T) {
	client := testutil.NewFixedClient(t)
	forex := &omise.Forex{}
	client.MustDo(forex, &RetrieveForex{
		Currency: "usd",
	})

	r.Equal(t, "USD", forex.Base)
	r.Equal(t, "THB", forex.Quote)
	r.Equal(t, float64(30.4847017), forex.Rate)
}

func TestForex_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	forex := &omise.Forex{}
	client.MustDo(forex, &RetrieveForex{
		Currency: "USD",
	})

	r.Equal(t, forex.Object, "forex")
	testutil.LogObj(t, forex)
}
