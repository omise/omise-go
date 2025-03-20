package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
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
			`{"amount":10000,"currency":"thb","platform_fee":{}}`,
		},
		{
			&CreateCharge{
				Amount:      10000,
				Currency:    "thb",
				DontCapture: true,
			},
			`{"amount":10000,"currency":"thb","platform_fee":{},"capture":false}`,
		},
		{
			&CreateCharge{
				Amount:            10000,
				Currency:          "thb",
				DontCapture:       true,
				AuthorizationType: omise.PreAuth,
			},
			`{"amount":10000,"currency":"thb","authorization_type":"pre_auth","platform_fee":{},"capture":false}`,
		},
		{
			&CreateCharge{
				Amount:            10000,
				Currency:          "thb",
				DontCapture:       true,
				AuthorizationType: omise.FinalAuth,
			},
			`{"amount":10000,"currency":"thb","authorization_type":"final_auth","platform_fee":{},"capture":false}`,
		},
		{
			&CreateCharge{
				Amount:           10000,
				Currency:         "thb",
				WebhookEndpoints: []string{"https://docs.opn.ooo/api-webhooks"},
			},
			`{"amount":10000,"currency":"thb","webhook_endpoints":["https://docs.opn.ooo/api-webhooks"],"platform_fee":{}}`,
		},
		{
			&CreateCharge{
				Amount:           10000,
				Currency:         "thb",
				WebhookEndpoints: []string{"https://docs.opn.ooo/api-webhooks"},
				Description:      "This is a card cahrge",
				Metadata: map[string]interface{}{
					"Hello": "World",
				},
				Ip:                   "192.168.1.1",
				TransactionIndicator: omise.MIT,
				RecurringReason:      omise.Blank,
				PlatformFee: omise.PlatformFee{
					Fixed:      10,
					Percentage: 2,
				},
			},
			`{"amount":10000,"currency":"thb","description":"This is a card cahrge","metadata":{"Hello":"World"},"webhook_endpoints":["https://docs.opn.ooo/api-webhooks"],"ip":"192.168.1.1","transaction_indicator":"MIT","platform_fee":{"fixed":10,"percentage":2}}`,
		},
	}
	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestRetrieveCharge_PayNow(t *testing.T) {
	const ChargeID = "chrg_test_5nvdpnapyle90ejudk8"
	client := testutil.NewFixedClient(t)

	charge := &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeID})

	r.Equal(t, ChargeID, charge.ID)
	r.NotNil(t, charge.Source.ScannableCode)
	r.Equal(t, "qrcode.png", charge.Source.ScannableCode.Image.Filename)
	r.Equal(t, "https://api.omise.co/charges/chrg_test_5nvdpnapyle90ejudk8/documents/docu_test_5nvdpnc12jbp0hwy031/downloads/542F307C03299822", charge.Source.ScannableCode.Image.DownloadURI)
}

func TestRetrieveCharge_BillPayment(t *testing.T) {
	const ChargeID = "chrg_test_5kdpo8ie121323la8ew"
	client := testutil.NewFixedClient(t)

	charge := &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeID})

	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, "bill_payment_tesco_lotus", charge.Source.Type)
	r.False(t, charge.Source.ZeroInterestInstallments)
	r.Equal(t, "https://api.omise.co/charges/chrg_test_5kdpo8ie121323la8ew/documents/docu_test_5kdpo8jcq8jn1zke3i1/downloads/FABFCC17CBD02176", charge.Source.References.Barcode)
}

func TestCharge(t *testing.T) {
	const (
		ChargeID               = "chrg_test_4yq7duw15p9hdrjp8oq"
		ChargeIdPartialCapture = "chrg_test_5x1753iuub61dfe41q4"
		TransactionID          = "trxn_test_4yq7duwb9jts1vxgqua"
		CardID                 = "card_test_4yq6tuucl9h4erukfl0"
		RefundID               = "rfnd_test_4yqmv79ahghsiz23y3c"
	)

	client := testutil.NewFixedClient(t)

	charge := &omise.Charge{}
	client.MustDo(charge, &CreateCharge{})
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, true, charge.Capture)
	r.Equal(t, omise.AuthorizationType(""), charge.AuthorizationType)
	r.Equal(t, int64(0), charge.CapturedAmount)
	r.Equal(t, int64(0), charge.AuthorizedAmount)

	//partial capture
	charge = &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeIdPartialCapture})
	r.Equal(t, ChargeIdPartialCapture, charge.ID)
	r.Equal(t, false, charge.Capture)
	r.Equal(t, omise.PreAuth, charge.AuthorizationType)
	r.Equal(t, int64(5000), charge.CapturedAmount)
	r.Equal(t, int64(10000), charge.AuthorizedAmount)
	r.Equal(t, int64(10000), charge.Amount)

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
		Amount:      409669,
		Currency:    "thb",
		DontCapture: true,
		Card:        token.ID,
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

	expected := `{"customer":"customer_id","amount":100000,"currency":"thb","metadata":{"color":"red"},"platform_fee":{}}`

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

	expected := `{"customer":"customer_id","amount":100000,"currency":"thb","platform_fee":{}}`

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

func TestRetrieveCharge_HasExpiredAt(t *testing.T) {
	const ChargeID = "chrg_test_5kdpo8ie121323la8ew"
	client := testutil.NewFixedClient(t)

	charge := &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeID})

	r.Equal(t, "2020-07-01 03:53:01 +0000 UTC", charge.ExpiresAt.String())
}

func TestRetrieveCharge_Has3DSFields(t *testing.T) {
	const ChargeID = "chrg_test_4yq7duw15p9hdrjp8oq"
	client := testutil.NewFixedClient(t)

	charge := &omise.Charge{}
	client.MustDo(charge, &RetrieveCharge{ChargeID})

	r.Equal(t, "additional authentication required", charge.MerchantAdvice)
	r.Equal(t, "03", charge.MerchantAdviceCode)
	r.Equal(t, []string{"phone_number", "address"}, charge.Missing3DSFields)
}
