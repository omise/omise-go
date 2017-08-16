package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestReceipt(t *testing.T) {
	const (
		ReceiptID = "rcpt_test_12345"
	)

	client := testutil.NewFixedClient(t)
	receipt := &omise.Receipt{}
	client.MustDo(receipt, &RetrieveReceipt{ReceiptID})
	a.Equal(t, ReceiptID, receipt.ID)
	if a.NotNil(t, receipt.Number) {
		a.Equal(t, "1", receipt.Number)
	}

	receipts := &omise.ReceiptList{}
	client.MustDo(receipts, &ListReceipts{})
	if a.Len(t, receipts.Data, 1) {
		a.Equal(t, ReceiptID, receipts.Data[0].ID)
	}
}
