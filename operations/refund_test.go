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

	charge, createCharge := &omise.Charge{}, &CreateCharge{
		Amount:      819229,
		Currency:    "thb",
		Description: "should be refunded soon.",
		Card:        token.ID,
	}
	client.MustDo(charge, createCharge)

	a.Equal(t, createCharge.Amount, charge.Amount)
	a.Equal(t, createCharge.Currency, charge.Currency)

	// list refunds on the charge
	refunds, list := &omise.RefundList{}, &ListRefunds{
		charge.ID,
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	client.MustDo(refunds, list)

	a.Len(t, refunds.Data, 0)

	// create a half refund on the charge
	refund, create := &omise.Refund{}, &CreateRefund{
		ChargeID: charge.ID,
		Amount:   charge.Amount >> 1,
	}
	client.MustDo(refund, create)

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
	retrieve := &RetrieveRefund{
		ChargeID: charge.ID,
		RefundID: refund.ID,
	}
	client.MustDo(refund, retrieve)

	a.Equal(t, refunds.Data[0].ID, refund.ID)
	a.Equal(t, refunds.Data[0].Amount, refund.Amount)
}
