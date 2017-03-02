package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	client := testutil.NewFixedClient(t)
	account := &omise.Account{}
	client.MustDo(account, &RetrieveAccount{})
	r.Equal(t, account.ID, "acct_4yq6tcsyoged5c0ocxd")
}

func TestAccount_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	account := &omise.Account{}
	client.MustDo(account, &RetrieveAccount{})
	r.Equal(t, account.Object, "account")

	testutil.LogObj(t, account)
}
