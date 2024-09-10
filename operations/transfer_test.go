package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	const (
		RecipientID = "recp_test_50894vc13y8z4v51iuc"
		TransferID  = "trsf_test_4yqacz8t3cbipcj766u"
	)

	client := testutil.NewFixedClient(t)

	transfer := &omise.Transfer{}
	client.MustDo(transfer, &CreateTransfer{Amount: 192188})
	r.Equal(t, TransferID, transfer.ID)
	r.Equal(t, int64(192188), transfer.Amount)

	transfer = &omise.Transfer{}
	client.MustDo(transfer, &RetrieveTransfer{TransferID})
	r.Equal(t, TransferID, transfer.ID)
	r.Equal(t, RecipientID, transfer.Recipient)
	r.NotNil(t, transfer.BankAccount)
	r.Equal(t, "6789", transfer.BankAccount.LastDigits)

	del := &omise.Deletion{}
	client.MustDo(del, &DestroyTransfer{TransferID})
	r.Equal(t, TransferID, del.ID)
	r.True(t, del.Deleted)
}

func TestTransfer_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// make a transfer to default recipient. (empty RecipientID)
	transfer := &omise.Transfer{}
	client.MustDo(transfer, &CreateTransfer{Amount: 32100})

	r.Equal(t, int64(32100), transfer.Amount)
	r.NotNil(t, transfer.BankAccount)

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

	r.True(t, len(transfers.Data) > 0, "no transfer was created.")

	transfer2 = transfers.Find(transfer.ID)
	r.Equal(t, transfer.ID, transfer2.ID)
	r.Equal(t, transfer.Amount, transfer2.Amount)

	// destroy transfer
	del, destroy := &omise.Deletion{}, &DestroyTransfer{TransferID: transfer.ID}
	client.MustDo(del, destroy)

	r.Equal(t, transfer.Object, del.Object)
	r.Equal(t, transfer.ID, del.ID)
	r.Equal(t, transfer.Live, del.Live)
	r.True(t, del.Deleted)
}

func TestCreateTransferMarshal_WithMetadata(t *testing.T) {
	req := &CreateTransfer{
		Amount: 192188,
		Metadata: map[string]interface{}{
			"color": "red",
		},
	}

	expected := `{"amount":192188,"metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestCreateTransferMarshal_WithoutMetadata(t *testing.T) {
	req := &CreateTransfer{
		Amount: 192188,
	}

	expected := `{"amount":192188}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}
