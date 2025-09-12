package operations_test

import (
	"encoding/json"
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	const TokenID = "tokn_test_4yq8lbecl0q6dsjzxr5"
	client := testutil.NewFixedClient(t)

	token := &omise.Token{}
	client.MustDo(token, &CreateToken{
		Name:            "JOHN DOE",
		Number:          "4242424242424242",
		ExpirationMonth: 1,
		ExpirationYear:  2017,
		SecurityCode:    "123",
	})
	r.Equal(t, TokenID, token.ID)

	token = &omise.Token{}
	client.MustDo(token, &RetrieveToken{TokenID})
	r.Equal(t, TokenID, token.ID)
}

func TestToken_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	tok1, tok2 := createTestToken(client), &omise.Token{}
	client.MustDo(tok2, &RetrieveToken{ID: tok1.ID})

	r.Equal(t, *tok1, *tok2)
}

func TestToken_WithEmailPhone_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	token := &omise.Token{}
	client.MustDo(token, &CreateToken{
		Name:            "John Doe",
		Number:          "4242424242424242",
		ExpirationMonth: 12,
		ExpirationYear:  2025,
		SecurityCode:    "123",
		Email:           "john@example.com",
		PhoneNumber:     "+66812345678",
		City:            "Bangkok",
		PostalCode:      "10240",
	})

	r.NotEmpty(t, token.ID)
}

func TestCreateToken_EmailPhone_Marshal(t *testing.T) {
	req := &CreateToken{
		Name:            "John Doe",
		Number:          "4242424242424242",
		ExpirationMonth: 12,
		ExpirationYear:  2025,
		SecurityCode:    "123",
		Email:           "john@example.com",
		PhoneNumber:     "+66812345678",
		City:            "Bangkok",
		PostalCode:      "10240",
	}

	b, err := json.Marshal(req)
	r.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(b, &result)
	r.NoError(t, err)

	card, ok := result["card"].(map[string]interface{})
	r.True(t, ok, "card field should exist")
	r.Equal(t, "john@example.com", card["email"])
	r.Equal(t, "+66812345678", card["phone_number"])
	r.Equal(t, "Bangkok", card["city"])
	r.Equal(t, "10240", card["postal_code"])
}
