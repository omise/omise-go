package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
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
	a.Equal(t, TokenID, token.ID)

	token = &omise.Token{}
	client.MustDo(token, &RetrieveToken{TokenID})
	a.Equal(t, TokenID, token.ID)
}

func TestToken_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	tok1, tok2 := createTestToken(client), &omise.Token{}
	client.MustDo(tok2, &RetrieveToken{ID: tok1.ID})

	a.Equal(t, *tok1, *tok2)
}
