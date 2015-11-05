package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	account := &omise.Account{}
	if e := client.Do(account, &RetrieveAccount{}); !a.NoError(t, e) {
		return
	}

	a.Equal(t, account.Object, "account")
	testutil.LogObj(t, account)
}
