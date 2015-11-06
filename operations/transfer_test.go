package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// make a transfer to default recipient. (empty RecipientID)
	transfer := &omise.Transfer{}
	client.MustDo(transfer, &CreateTransfer{Amount: 32100})

	a.Equal(t, int64(32100), transfer.Amount)
	a.NotNil(t, transfer.BankAccount)

	// gets created transfer
	transfer2 := &omise.Transfer{}
	client.MustDo(transfer2, &RetrieveTransfer{
		TransferID: transfer.ID,
	})

	// list transfers
	transfers := &omise.TransferList{}
	client.MustDo(transfers, &ListTransfers{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	})

	a.True(t, len(transfers.Data) > 0, "no transfer was created.")

	transfer2 = transfers.Find(transfer.ID)
	a.Equal(t, transfer.ID, transfer2.ID)
	a.Equal(t, transfer.Amount, transfer2.Amount)

	// update transfer
	transfer2 = &omise.Transfer{}
	client.MustDo(transfer2, &UpdateTransfer{
		TransferID: transfer.ID,
		Amount:     12300,
	})

	a.Equal(t, transfer.ID, transfer2.ID)
	a.Equal(t, int64(12300), transfer2.Amount)

	// destroy transfer
	del, destroy := &omise.Deletion{}, &DestroyTransfer{TransferID: transfer.ID}
	client.MustDo(del, destroy)

	a.Equal(t, transfer.Object, del.Object)
	a.Equal(t, transfer.ID, del.ID)
	a.Equal(t, transfer.Live, del.Live)
	a.True(t, del.Deleted)
}
