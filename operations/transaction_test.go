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
	const (
		TransactionID  = "trxn_test_4yq7duwb9jts1vxgqua"
		TransactionID2 = "trxn_test_4yqafnvlztbf3908vs1"
	)

	client := testutil.NewFixedClient(t)

	tx := &omise.Transaction{}
	client.MustDo(tx, &RetrieveTransaction{TransactionID})
	a.Equal(t, TransactionID, tx.ID)

	transactions := &omise.TransactionList{}
	client.MustDo(transactions, &ListTransactions{})
	if a.Len(t, transactions.Data, 2) {
		a.Equal(t, TransactionID, transactions.Data[0].ID)
		a.Equal(t, TransactionID2, transactions.Data[1].ID)
	}
}

func TestTransaction_Network(t *testing.T) {
	// TODO: There is no way to approve a recipient and verify transfer programmatically so
	//   we can't create transaction programmtically, yet. Thus this is skipped, for now.
	//   You can try this test by manually completing one transfer on the test dashboard
	//   which will cause a `Transaction` object to be created.
	t.Skip()

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// list transactions
	transactions := &omise.TransactionList{}
	client.MustDo(transactions, &ListTransactions{
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
