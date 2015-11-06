package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestRefund(t *testing.T) {
	client := testutil.NewTestClient(t)

	// create a charge so we can refund
	// TODO: DRY with TestCharge
	token := &omise.Token{}
	client.MustDo(token, CreateTokenOp)

	charge := &omise.Charge{}
	client.MustDo(charge, &CreateCharge{
		Amount:      819229,
		Currency:    "thb",
		Description: "should be refunded soon.",
		Card:        token.ID,
	})

	a.Equal(t, int64(819229), charge.Amount)
	a.Equal(t, "thb", charge.Currency)

	// list refunds on the charge
	list := &ListRefunds{
		charge.ID,
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}

	refunds := &omise.RefundList{}
	client.MustDo(refunds, list)

	a.Len(t, refunds.Data, 0)

	// create a half refund on the charge
	refund := &omise.Refund{}
	client.MustDo(refund, &CreateRefund{
		ChargeID: charge.ID,
		Amount:   charge.Amount >> 1,
	})

	a.Equal(t, refund.Amount, charge.Amount>>1)
	a.Equal(t, refund.Currency, charge.Currency)

	// list refunds again which now contains the created refunds
	client.MustDo(refunds, list)

	if a.Len(t, refunds.Data, 1) {
		a.Equal(t, refunds.Data[0].ID, refund.ID)
	} else {
		return
	}

	// retrieve refund by id, which should match what we already have.
	client.MustDo(refund, &RetrieveRefund{
		ChargeID: charge.ID,
		RefundID: refund.ID,
	})

	a.Equal(t, refunds.Data[0].ID, refund.ID)
	a.Equal(t, refunds.Data[0].Amount, refund.Amount)
}
