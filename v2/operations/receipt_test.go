package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	a "github.com/stretchr/testify/assert"
)

func TestReceipt(t *testing.T) {
	const (
		ReceiptID = "rcpt_test_12345"
	)

	client := testutil.NewFixedClient(t)
	receipt, _ := client.Receipt().Retrieve(context.Background(), &omise.RetrieveReceiptParams{ReceiptID})
	a.Equal(t, ReceiptID, receipt.ID)
	if a.NotNil(t, receipt.Number) {
		a.Equal(t, "1", receipt.Number)
	}

	receipts := &omise.ReceiptList{}
	receipts, _ = client.Receipt().List(context.Background(), &omise.ListReceiptsParams{})
	if a.Len(t, receipts.Data, 1) {
		a.Equal(t, ReceiptID, receipts.Data[0].ID)
	}
}
