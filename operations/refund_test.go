package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
)

func TestRefund(t *testing.T) {
	const (
		ChargeID      = "chrg_test_4yq7duw15p9hdrjp8oq"
		TransactionID = "trxn_test_4yqmv79fzpy0gmz5mmq"
		RefundID      = "rfnd_test_4yqmv79ahghsiz23y3c"
	)

	client := testutil.NewFixedClient(t)

	refund := &omise.Refund{}
	client.MustDo(refund, &RetrieveRefund{ChargeID, RefundID})
	r.Equal(t, RefundID, refund.ID)
	r.Equal(t, ChargeID, refund.Charge)
	r.Equal(t, TransactionID, refund.Transaction)

	refund = &omise.Refund{}
	client.MustDo(refund, &CreateRefund{ChargeID, 10000, false})
	r.Equal(t, RefundID, refund.ID)
	r.Equal(t, int64(10000), refund.Amount)

	e := client.Do(nil, &RetrieveRefund{ChargeID, "not_exist"})
	r.Error(t, e)
	r.EqualError(t, e, "(404/not_found) refund 404 was not found")
}

func TestRefund_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// create a charge so we can refund
	token := createTestToken(client)
	charge := createTestCharge(client, token)

	// list refunds on the charge
	list := &ListRefunds{
		charge.ID,
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}

	refunds := &omise.RefundList{}
	client.MustDo(refunds, list)

	r.Len(t, refunds.Data, 0)

	// create a half refund on the charge
	refund := &omise.Refund{}
	client.MustDo(refund, &CreateRefund{
		ChargeID: charge.ID,
		Amount:   charge.Amount >> 1,
	})

	r.Equal(t, refund.Amount, charge.Amount>>1)
	r.Equal(t, refund.Currency, charge.Currency)

	// list refunds again which now contains the created refunds
	client.MustDo(refunds, list)

	r.Len(t, refunds.Data, 1)
	r.Equal(t, refunds.Data[0].ID, refund.ID)

	// retrieve refund by id, which should match what we already have.
	client.MustDo(refund, &RetrieveRefund{
		ChargeID: charge.ID,
		RefundID: refund.ID,
	})

	r.Equal(t, refunds.Data[0].ID, refund.ID)
	r.Equal(t, refunds.Data[0].Amount, refund.Amount)
}
