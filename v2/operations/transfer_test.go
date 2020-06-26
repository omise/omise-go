package operations_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	const (
		RecipientID = "recp_test_50894vc13y8z4v51iuc"
		TransferID  = "trsf_test_4yqacz8t3cbipcj766u"
	)

	client := testutil.NewFixedClient(t)

	transfer, _ := client.Transfer().Create(context.Background(), &omise.CreateTransferParams{Amount: 192188})
	r.Equal(t, TransferID, transfer.ID)
	r.Equal(t, int64(192188), transfer.Amount)

	transfer, _ = client.Transfer().Retrieve(context.Background(), &omise.RetrieveTransferParams{TransferID})
	r.Equal(t, TransferID, transfer.ID)
	r.Equal(t, RecipientID, transfer.Recipient)
	r.NotNil(t, transfer.BankAccount)
	r.Equal(t, "6789", transfer.BankAccount.LastDigits)

	transfer, _ = client.Transfer().Update(context.Background(), &omise.UpdateTransferParams{
		TransferID: TransferID,
		Amount:     192189,
	})
	r.Equal(t, TransferID, transfer.ID)
	r.Equal(t, int64(192189), transfer.Amount)

	del, _ := client.Transfer().Destroy(context.Background(), &omise.DestroyTransferParams{TransferID})
	r.Equal(t, TransferID, del.ID)
}

func TestTransfer_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// make a transfer to default recipient. (empty RecipientID)
	transfer, _ := client.Transfer().Create(context.Background(), &omise.CreateTransferParams{Amount: 32100})

	r.Equal(t, int64(32100), transfer.Amount)
	r.NotNil(t, transfer.BankAccount)

	// gets created transfer
	transfer2, _ := client.Transfer().Retrieve(context.Background(), &omise.RetrieveTransferParams{
		TransferID: transfer.ID,
	})

	// list transfers
	transfers, _ := client.Transfer().List(context.Background(), &omise.ListTransfersParams{
		omise.ListParams{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	})

	r.True(t, len(transfers.Data) > 0, "no transfer was created.")

	transfer2 = transfers.Find(transfer.ID)
	r.Equal(t, transfer.ID, transfer2.ID)
	r.Equal(t, transfer.Amount, transfer2.Amount)

	// update transfer
	transfer2, _ = client.Transfer().Update(context.Background(), &omise.UpdateTransferParams{
		TransferID: transfer.ID,
		Amount:     12300,
	})

	r.Equal(t, transfer.ID, transfer2.ID)
	r.Equal(t, int64(12300), transfer2.Amount)

	// destroy transfer
	del, destroy := &omise.Deletion{}, &DestroyTransfer{TransferID: transfer.ID}
	client.MustDo(del, destroy)

	r.Equal(t, transfer.Object, del.Object)
	r.Equal(t, transfer.ID, del.ID)
	r.Equal(t, transfer.Live, del.Live)
}

func TestCreateTransferMarshal_WithMetadata(t *testing.T) {
	req := &omise.CreateTransferParams{
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
	req := &omise.CreateTransferParams{
		Amount: 192188,
	}

	expected := `{"amount":192188}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateTransferMarshal_WithMetadata(t *testing.T) {
	req := &omise.UpdateTransferParams{
		TransferID: "trsf_test_4yqacz8t3cbipcj766u",
		Amount:     192188,
		Metadata: map[string]interface{}{
			"color": "red",
		},
	}

	expected := `{"amount":192188,"metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateTransferMarshal_WithoutMetadata(t *testing.T) {
	req := &omise.UpdateTransferParams{
		TransferID: "trsf_test_4yqacz8t3cbipcj766u",
		Amount:     192188,
	}

	expected := `{"amount":192188}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}
