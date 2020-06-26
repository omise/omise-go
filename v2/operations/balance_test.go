package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	client := testutil.NewFixedClient(t)

	balance, err := client.Balance().Retrieve(context.Background(), &omise.RetrieveBalanceParams{})

	r.Nil(t, err)
	r.Equal(t, int64(96094), balance.Total)
	r.Equal(t, "thb", balance.Currency)
}

func TestBalance_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	balance, err := client.Balance().Retrieve(context.Background(), &omise.RetrieveBalanceParams{})

	r.Nil(t, err)
	r.Equal(t, balance.Object, "balance")
	testutil.LogObj(t, balance)
}
