package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestCharge(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	token := &omise.Token{}
	client.MustDo(token, CreateTokenOp)

	// create
	charge, create := &omise.Charge{}, &CreateCharge{
		Amount:      204842,
		Currency:    "thb",
		Description: "initial charge.",
		Card:        token.ID,
	}
	client.MustDo(charge, create)

	a.Equal(t, create.Amount, charge.Amount)
	a.Equal(t, create.Currency, charge.Currency)

	// retrieve created charge
	charge2 := &omise.Charge{}
	client.MustDo(charge2, &RetrieveCharge{ChargeID: charge.ID})

	a.Equal(t, charge.ID, charge2.ID)
	a.Equal(t, charge.Amount, charge2.Amount)
	a.Equal(t, charge.Description, charge2.Description)

	// list created charges from the last hour
	charges, list := &omise.ChargeList{}, &ListCharges{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	client.MustDo(&charges, list)

	a.True(t, len(charges.Data) > 0, "charges list empty!")
	charge2 = charges.Find(charge.ID)
	if !a.NotNil(t, charge2, "could not find recent charges in list.") {
		return
	}

	a.Equal(t, charge.ID, charge2.ID, "charge not in returned list.")
	a.Equal(t, charge.Amount, charge2.Amount, "listed charge has wrong amount.")

	// update charge
	charge2 = &omise.Charge{}
	update := &UpdateCharge{
		ChargeID:    charge.ID,
		Description: "updated charge.",
	}
	client.MustDo(charge2, update)

	a.Equal(t, charge.ID, charge2.ID)
	if a.NotNil(t, charge2.Description) {
		a.Equal(t, update.Description, *charge2.Description)
	}
}

func TestCharge_Uncaptured(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	token := &omise.Token{}
	client.MustDo(token, CreateTokenOp)

	// create uncaptured charge
	charge, create := &omise.Charge{}, &CreateCharge{
		Amount:      409669,
		Currency:    "thb",
		DontCapture: true,
		Card:        token.ID,
	}
	client.MustDo(charge, create)

	a.Equal(t, create.Amount, charge.Amount)
	a.False(t, charge.Captured, "charge unintentionally captured!")

	// then capture it
	charge2 := &omise.Charge{}
	client.MustDo(charge2, &CaptureCharge{ChargeID: charge.ID})

	a.Equal(t, charge.ID, charge2.ID)
	a.True(t, charge2.Captured, "charge not captured!")
}

func TestCharge_Invalid(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	token := &omise.Token{}
	client.MustDo(token, CreateTokenOp)

	e := client.Do(nil, &CreateCharge{
		Amount:   12345,
		Currency: "omd", // OMISE DOLLAR, why not?
		Card:     token.ID,
	})
	a.EqualError(t, e, "(400/invalid_charge) currency is currently not supported")

	e = client.Do(nil, &CreateCharge{
		Amount:   12345,
		Currency: "thb", // OMISE DOLLAR, why not?
		Card:     "tok_asdf",
	})
	a.EqualError(t, e, "(404/not_found) token tok_asdf was not found")
}
