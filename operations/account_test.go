package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestAccount_Network(t *testing.T) {
	client := testutil.NewTestClient(t)

	account := &omise.Account{}
	client.MustDo(account, &RetrieveAccount{})
	a.Equal(t, account.Object, "account")

	testutil.LogObj(t, account)
}
