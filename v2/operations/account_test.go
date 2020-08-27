package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	client := testutil.NewFixedClient(t)
	account, err := client.Account().Retrieve(context.Background(), &omise.RetrieveAccountParams{})

	r.Nil(t, err)
	r.Equal(t, account.ID, "acct_4yq6tcsyoged5c0ocxd")
}

func TestAccount_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	account, err := client.Account().Retrieve(context.Background(), &omise.RetrieveAccountParams{})

	r.Nil(t, err)
	r.Equal(t, account.Object, "account")

	testutil.LogObj(t, account)
}
