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
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	// make a transfer to default recipient. (empty RecipientID)
	transfer, create := &omise.Transfer{}, &CreateTransfer{
		Amount: 32100,
	}
	if e := client.Do(transfer, create); !a.NoError(t, e) {
		return
	}

	a.Equal(t, create.Amount, transfer.Amount)
	a.NotNil(t, transfer.BankAccount)

	// gets created transfer
	transfer2, retrieve := &omise.Transfer{}, &RetrieveTransfer{
		TransferID: transfer.ID,
	}
	if e := client.Do(transfer2, retrieve); !a.NoError(t, e) {
		return
	}

	// list transfers
	transfers, list := &omise.TransferList{}, &ListTransfers{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	if e := client.Do(transfers, list); !a.NoError(t, e) {
		return
	}

	a.True(t, len(transfers.Data) > 0, "no transfer was created.")

	transfer2 = transfers.Find(transfer.ID)
	a.Equal(t, transfer.ID, transfer2.ID)
	a.Equal(t, transfer.Amount, transfer2.Amount)

	// update transfer
	transfer2, update := &omise.Transfer{}, &UpdateTransfer{
		TransferID: transfer.ID,
		Amount:     12300,
	}
	if e := client.Do(transfer2, update); !a.NoError(t, e) {
		return
	}

	a.Equal(t, transfer.ID, transfer2.ID)
	a.Equal(t, update.Amount, transfer2.Amount)

	// destroy transfer
	del, destroy := &omise.Deletion{}, &DestroyTransfer{TransferID: transfer.ID}
	if e := client.Do(del, destroy); !a.NoError(t, e) {
		return
	}

	a.Equal(t, transfer.Object, del.Object)
	a.Equal(t, transfer.ID, del.ID)
	a.Equal(t, transfer.Live, del.Live)
	a.True(t, del.Deleted)
}
