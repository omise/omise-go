package omise_test

import (
	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
	"testing"
)

// TODO: Include token retrieval and customer creation into the test flow.
const (
	TestCustomer = "cust_test_51w32b7kmr04n6bkcp8"
	TestCard     = "card_test_51w32k09s2eu8azi12r"
)

func TestNewClient(t *testing.T) {
	pkey, skey := testKeys()

	if _, e := NewClient(pkey, skey); !a.NoError(t, e) {
		return
	}
	if _, e := NewClient("", skey); !a.NoError(t, e) {
		return
	}
	if _, e := NewClient(pkey, ""); !a.NoError(t, e) {
		return
	}

	if _, e := NewClient("", ""); a.Error(t, e) {
		a.Equal(t, ErrInvalidKey, e)
	}
	if _, e := NewClient("123", "123"); a.Error(t, e) {
		a.Equal(t, ErrInvalidKey, e)
	}
}

func TestClient_Charge(t *testing.T) {
	client, e := NewClient(testKeys())
	if !a.NoError(t, e) {
		return
	}

	op := &operations.CreateCharge{
		Amount:   204842,
		Currency: "thb",
		Customer: TestCustomer,
		Card:     TestCard,
	}

	charge, e := client.CreateCharge(op)
	if !(a.NoError(t, e) && a.NotNil(t, charge)) {
		return
	}

	a.Equal(t, op.Amount, charge.Amount)
	a.Equal(t, op.Currency, charge.Currency)
}

func TestClient_InvalidCharge(t *testing.T) {
	client, e := NewClient(testKeys())
	if !a.NoError(t, e) {
		return
	}

	_, e = client.CreateCharge(&operations.CreateCharge{
		Amount:   12345,
		Currency: "omd", // OMISE DOLLAR, why not?
		Customer: TestCustomer,
		Card:     TestCard,
	})
	a.EqualError(t, e, "(400/invalid_charge) currency is currently not supported")
}
