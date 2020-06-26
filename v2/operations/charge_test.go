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

func TestCreateChargeMarshal(t *testing.T) {
	testdata := []struct {
		req      *CreateCharge
		expected string
	}{
		{
			&CreateCharge{
				Amount:   10000,
				Currency: "THB",
			},
			`{"amount":10000,"currency":"THB"}`,
		},
		{
			&CreateCharge{
				Amount:   10000,
				Currency: "THB",
				Capture:  new(bool),
			},
			`{"amount":10000,"capture":false,"currency":"THB"}`,
		},
	}
	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestCharge(t *testing.T) {
	const (
		ChargeID      = "chrg_test_4yq7duw15p9hdrjp8oq"
		TransactionID = "trxn_test_4yq7duwb9jts1vxgqua"
		CardID        = "card_test_4yq6tuucl9h4erukfl0"
		RefundID      = "rfnd_test_4yqmv79ahghsiz23y3c"
	)

	client := testutil.NewFixedClient(t)

	charge, err := client.Charge().Create(context.Background(), &omise.CreateChargeParams{})
	r.Nil(t, err)
	r.Equal(t, ChargeID, charge.ID)

	charge, err = client.Charge().Retrieve(context.Background(), &omise.RetrieveChargeParams{ChargeID})
	r.Nil(t, err)
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, TransactionID, charge.Transaction)
	r.Equal(t, CardID, charge.Card.ID)
	r.Len(t, charge.Refunds.Data, 1)
	r.Equal(t, RefundID, charge.Refunds.Data[0].ID)

	charges, err := client.Charge().List(context.Background(), &omise.ListChargesParams{})
	r.Len(t, charges.Data, 1)
	r.Equal(t, "chrg_test_4yq7duw15p9hdrjp8oq", charges.Data[0].ID)

	charge, err = client.Charge().Update(context.Background(), &omise.UpdateChargeParams{
		ChargeID:    ChargeID,
		Description: "Charge for order 3947 (XXL)",
	})
	r.NotNil(t, charge.Description)
	r.Equal(t, "Charge for order 3947 (XXL)", *charge.Description)

	_, err = client.Charge().Retrieve(context.Background(), &omise.RetrieveChargeParams{"not_exist"})
	r.Error(t, err)
	r.EqualError(t, err, "(404/not_found) customer missing was not found")
}

func TestCharge_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create
	create := &omise.CreateChargeParams{
		Amount:      204842,
		Currency:    "THB",
		Description: "initial charge.",
		Card:        token.ID,
	}
	charge, err := client.Charge().Create(context.Background(), create)

	r.Equal(t, create.Amount, charge.Amount)
	r.Equal(t, create.Currency, charge.Currency)

	// retrieve created charge
	charge2, err := client.Charge().Retrieve(context.Background(), &omise.RetrieveChargeParams{ChargeID: charge.ID})

	r.Nil(t, err)
	r.Equal(t, charge.ID, charge2.ID)
	r.Equal(t, charge.Amount, charge2.Amount)
	r.Equal(t, charge.Description, charge2.Description)

	// list created charges from the last hour
	list := &omise.ListChargesParams{
		omise.ListParams{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	charges, _ := client.Charge().List(context.Background(), list)
	r.True(t, len(charges.Data) > 0, "charges list empty!")

	charge2 = charges.Find(charge.ID)
	r.NotNil(t, charge2, "could not find recent charges in list.")
	r.Equal(t, charge.ID, charge2.ID, "charge not in returned list.")
	r.Equal(t, charge.Amount, charge2.Amount, "listed charge has wrong amount.")

	// update charge
	update := &omise.UpdateChargeParams{
		ChargeID:    charge.ID,
		Description: "updated charge.",
	}
	charge2, _ = client.Charge().Update(context.Background(), update)

	r.Equal(t, charge.ID, charge2.ID)
	r.NotNil(t, charge2.Description)
	r.Equal(t, update.Description, *charge2.Description)
}

func TestCharge_Network_Uncaptured(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create uncaptured charge
	create := &omise.CreateChargeParams{
		Amount:   409669,
		Currency: "THB",
		Capture:  new(bool),
		Card:     token.ID,
	}
	charge, _ := client.Charge().Create(context.Background(), create)

	r.Equal(t, create.Amount, charge.Amount)
	r.False(t, charge.Paid, "charge unintentionally captured!")

	// then capture it
	charge2, err := client.Charge().Capture(context.Background(), &omise.CaptureChargeParams{ChargeID: charge.ID})

	r.Nil(t, err)
	r.Equal(t, charge.ID, charge2.ID)
	r.True(t, charge2.Paid, "charge not captured!")
}

func TestCharge_Network_Invalid(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	_, err := client.Charge().Create(context.Background(), &omise.CreateChargeParams{
		Amount:   12345,
		Currency: "omd", // OMISE DOLLAR, why not?
		Card:     token.ID,
	})
	r.EqualError(t, err, "(400/invalid_charge) currency is currently not supported")

	_, err = client.Charge().Create(context.Background(), &omise.CreateChargeParams{
		Amount:   12345,
		Currency: "THB",
		Card:     "tok_asdf",
	})
	r.EqualError(t, err, "(404/not_found) token tok_asdf was not found")
}

func TestCreateChargeMarshal_WithMetadata(t *testing.T) {
	req := &omise.CreateChargeParams{
		Customer: "customer_id",
		Amount:   100000,
		Currency: "THB",
		Metadata: map[string]interface{}{
			"color": "red",
		},
	}

	expected := `{"amount":100000,"currency":"THB","customer":"customer_id","metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "error should be nothing")
	r.Equal(t, expected, string(b))
}

func TestCreateChargeMarshal_WithoutMetadata(t *testing.T) {
	req := &omise.CreateChargeParams{
		Customer: "customer_id",
		Amount:   100000,
		Currency: "THB",
	}

	expected := `{"amount":100000,"currency":"THB","customer":"customer_id"}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateChargeMarshal_WithMetadata(t *testing.T) {
	req := &omise.UpdateChargeParams{
		ChargeID:    "chrg_test_4yq7duw15p9hdrjp8oq",
		Description: "Charge for order 3947 (XXL)",
		Metadata: map[string]interface{}{
			"color": "red",
		},
	}

	expected := `{"description":"Charge for order 3947 (XXL)","metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}
