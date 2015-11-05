package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	balance := &omise.Balance{}
	if e := client.Do(balance, &RetrieveBalance{}); !a.NoError(t, e) {
		return
	}

	a.Equal(t, balance.Object, "balance")
	testutil.LogObj(t, balance)
}
