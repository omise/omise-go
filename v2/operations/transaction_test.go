package operations_test

import (
	"context"
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	const (
		TransactionID  = "trxn_test_4yq7duwb9jts1vxgqua"
		TransactionID2 = "trxn_test_4yqafnvlztbf3908vs1"
	)

	client := testutil.NewFixedClient(t)

	tx, _ := client.Transaction().Retrieve(context.Background(), &omise.RetrieveTransactionParams{TransactionID})
	r.Equal(t, TransactionID, tx.ID)

	transactions, _ := client.Transaction().List(context.Background(), &omise.ListTransactionsParams{})
	r.Len(t, transactions.Data, 2)
	r.Equal(t, TransactionID, transactions.Data[0].ID)
	r.Equal(t, TransactionID2, transactions.Data[1].ID)
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
	transactions, _ := client.Transaction().List(context.Background(), &omise.ListTransactionsParams{
		omise.ListParams{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	})

	r.True(t, len(transactions.Data) > 0, "no transactions was created!")

	// retrieve a sample transaction
	transaction, _ := client.Transaction().Retrieve(context.Background(), &omise.RetrieveTransactionParams{
		TransactionID: transactions.Data[0].ID,
	})

	r.Equal(t, transactions.Data[0].ID, transaction.ID)
	r.Equal(t, transactions.Data[0].Amount, transaction.Amount)
}
