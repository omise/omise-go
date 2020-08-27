package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/assert"
)

func TestForex(t *testing.T) {
	client := testutil.NewFixedClient(t)
	forex, _ := client.Forex().Retrieve(context.Background(), &omise.RetrieveForexParams{
		Currency: "usd",
	})

	r.Equal(t, "USD", forex.Base)
	r.Equal(t, "THB", forex.Quote)
	r.Equal(t, float64(30.4847017), forex.Rate)
}

func TestForex_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	forex, _ := client.Forex().Retrieve(context.Background(), &omise.RetrieveForexParams{
		Currency: "USD",
	})

	r.Equal(t, forex.Object, "forex")
	testutil.LogObj(t, forex)
}
