package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	client := testutil.NewFixedClient(t)
	balance := &omise.Balance{}
	client.MustDo(balance, &RetrieveBalance{})
	a.Equal(t, int64(96094), balance.Total)
	a.Equal(t, "thb", balance.Currency)
}

func TestBalance_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	balance := &omise.Balance{}
	client.MustDo(balance, &RetrieveBalance{})

	a.Equal(t, balance.Object, "balance")
	testutil.LogObj(t, balance)
}
