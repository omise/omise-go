package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	// TODO: There is no way to approve a recipient and verify transfer programmatically so
	//   we can't create transaction programmtically, yet. Thus this is skipped, for now.
	//   You can try this test by manually completing one transfer on the test dashboard
	//   which will cause a `Transaction` object to be created.
	t.Skip()

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// list transactions
	transactions := &omise.TransactionList{}
	client.MustDo(transactions, &ListTransaction{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	})

	if !a.True(t, len(transactions.Data) > 0, "no transactions was created!") {
		return
	}

	// retrieve a sample transaction
	transaction := &omise.Transaction{}
	client.MustDo(transaction, &RetrieveTransaction{
		TransactionID: transactions.Data[0].ID,
	})

	a.Equal(t, transactions.Data[0].ID, transaction.ID)
	a.Equal(t, transactions.Data[0].Amount, transaction.Amount)
}
