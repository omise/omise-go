package operations_test

import (
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
				Currency: "thb",
			},
			`{"amount":10000,"currency":"thb"}`,
		},
		{
			&CreateCharge{
				Amount:   10000,
				Currency: "thb",
				Capture:  new(bool),
			},
			`{"amount":10000,"capture":false,"currency":"thb"}`,
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

	charge := &omise.Charge{}
	client.MustDo(charge, &CreateCharge{})
	r.Equal(t, ChargeID, charge.ID)

	charge = &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeID})
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, TransactionID, charge.Transaction)
	r.Equal(t, CardID, charge.Card.ID)
	r.Len(t, charge.Refunds.Data, 1)
	r.Equal(t, RefundID, charge.Refunds.Data[0].ID)

	charges := &omise.ChargeList{}
	client.MustDo(charges, &ListCharges{})
	r.Len(t, charges.Data, 1)
	r.Equal(t, "chrg_test_4yq7duw15p9hdrjp8oq", charges.Data[0].ID)

	client.MustDo(charge, &UpdateCharge{
		ChargeID:    ChargeID,
		Description: "Charge for order 3947 (XXL)",
	})
	r.NotNil(t, charge.Description)
	r.Equal(t, "Charge for order 3947 (XXL)", *charge.Description)

	err := client.Do(nil, &RetrieveCharge{"not_exist"})
	r.Error(t, err)
	r.EqualError(t, err, "(404/not_found) customer missing was not found")
}

func TestCharge_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create
	charge, create := &omise.Charge{}, &CreateCharge{
		Amount:      204842,
		Currency:    "thb",
		Description: "initial charge.",
		Card:        token.ID,
	}
	client.MustDo(charge, create)

	r.Equal(t, create.Amount, charge.Amount)
	r.Equal(t, create.Currency, charge.Currency)

	// retrieve created charge
	charge2 := &omise.Charge{}
	client.MustDo(charge2, &RetrieveCharge{ChargeID: charge.ID})

	r.Equal(t, charge.ID, charge2.ID)
	r.Equal(t, charge.Amount, charge2.Amount)
	r.Equal(t, charge.Description, charge2.Description)

	// list created charges from the last hour
	charges, list := &omise.ChargeList{}, &ListCharges{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	client.MustDo(&charges, list)
	r.True(t, len(charges.Data) > 0, "charges list empty!")

	charge2 = charges.Find(charge.ID)
	r.NotNil(t, charge2, "could not find recent charges in list.")
	r.Equal(t, charge.ID, charge2.ID, "charge not in returned list.")
	r.Equal(t, charge.Amount, charge2.Amount, "listed charge has wrong amount.")

	// update charge
	charge2 = &omise.Charge{}
	update := &UpdateCharge{
		ChargeID:    charge.ID,
		Description: "updated charge.",
	}
	client.MustDo(charge2, update)

	r.Equal(t, charge.ID, charge2.ID)
	r.NotNil(t, charge2.Description)
	r.Equal(t, update.Description, *charge2.Description)
}

func TestCharge_Network_Uncaptured(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create uncaptured charge
	charge, create := &omise.Charge{}, &CreateCharge{
		Amount:   409669,
		Currency: "thb",
		Capture:  new(bool),
		Card:     token.ID,
	}
	client.MustDo(charge, create)

	r.Equal(t, create.Amount, charge.Amount)
	r.False(t, charge.Paid, "charge unintentionally captured!")

	// then capture it
	charge2 := &omise.Charge{}
	client.MustDo(charge2, &CaptureCharge{ChargeID: charge.ID})

	r.Equal(t, charge.ID, charge2.ID)
	r.True(t, charge2.Paid, "charge not captured!")
}

func TestCharge_Network_Invalid(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	err := client.Do(nil, &CreateCharge{
		Amount:   12345,
		Currency: "omd", // OMISE DOLLAR, why not?
		Card:     token.ID,
	})
	r.EqualError(t, err, "(400/invalid_charge) currency is currently not supported")

	err = client.Do(nil, &CreateCharge{
		Amount:   12345,
		Currency: "thb",
		Card:     "tok_asdf",
	})
	r.EqualError(t, err, "(404/not_found) token tok_asdf was not found")
}

func TestCreateChargeMarshal_WithMetadata(t *testing.T) {
	req := &CreateCharge{
		Customer: "customer_id",
		Amount:   100000,
		Currency: "thb",
		Metadata: map[string]interface{}{
			"color": "red",
		},
	}

	expected := `{"amount":100000,"currency":"thb","customer":"customer_id","metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "error should be nothing")
	r.Equal(t, expected, string(b))
}

func TestCreateChargeMarshal_WithoutMetadata(t *testing.T) {
	req := &CreateCharge{
		Customer: "customer_id",
		Amount:   100000,
		Currency: "thb",
	}

	expected := `{"amount":100000,"currency":"thb","customer":"customer_id"}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateChargeMarshal_WithMetadata(t *testing.T) {
	req := &UpdateCharge{
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
