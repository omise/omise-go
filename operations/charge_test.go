package operations_test

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestCharge(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	charge, token := &omise.Charge{}, &omise.Token{}
	if e := client.Do(token, CreateTokenOp); !a.NoError(t, e) {
		return
	}

	createCharge := &CreateCharge{
		Amount:   204842,
		Currency: "thb",
		Card:     token.ID,
	}
	if e := client.Do(charge, createCharge); !(a.NoError(t, e) && a.NotNil(t, charge)) {
		return
	}

	a.Equal(t, createCharge.Amount, charge.Amount)
	a.Equal(t, createCharge.Currency, charge.Currency)
}

func TestCharge_Invalid(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	token := &omise.Token{}
	if e := client.Do(token, CreateTokenOp); !a.NoError(t, e) {
		return
	}

	e = client.Do(nil, &CreateCharge{
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
