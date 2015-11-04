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
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	// create a charge so we can refund
	// TODO: DRY with TestCharge
	token := &omise.Token{}
	if e := client.Do(token, CreateTokenOp); !a.NoError(t, e) {
		return
	}

	charge, createCharge := &omise.Charge{}, &CreateCharge{
		Amount:      819229,
		Currency:    "thb",
		Description: "should be refunded soon.",
		Card:        token.ID,
	}
	if e := client.Do(charge, createCharge); !(a.NoError(t, e) && a.NotNil(t, charge)) {
		return
	}

	a.Equal(t, createCharge.Amount, charge.Amount)
	a.Equal(t, createCharge.Currency, charge.Currency)

	// list refunds on the charge
	refunds, list := &omise.RefundList{}, &ListRefunds{
		charge.ID,
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	if e := client.Do(refunds, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, refunds.Data, 0)

	// create a half refund on the charge
	refund, create := &omise.Refund{}, &CreateRefund{
		ChargeID: charge.ID,
		Amount:   charge.Amount >> 1,
	}
	if e := client.Do(refund, create); !a.NoError(t, e) {
		return
	}

	a.Equal(t, refund.Amount, charge.Amount>>1)
	a.Equal(t, refund.Currency, charge.Currency)

	// list refunds again which now contains the created refunds
	if e := client.Do(refunds, list); !a.NoError(t, e) {
		return
	}

	if a.Len(t, refunds.Data, 1) {
		a.Equal(t, refunds.Data[0].ID, refund.ID)
	} else {
		return
	}

	// retreive refund by id, which should match what we already have.
	retreive := &RetreiveRefund{
		ChargeID: charge.ID,
		RefundID: refund.ID,
	}
	if e := client.Do(refund, retreive); !a.NoError(t, e) {
		return
	}

	a.Equal(t, refunds.Data[0].ID, refund.ID)
	a.Equal(t, refunds.Data[0].Amount, refund.Amount)
}
