package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	"github.com/omise/omise-go/v2/operations"
)

func TestMain(m *testing.M) {
	testutil.Main(m)
}

func createTestToken(client *testutil.TestClient) *omise.Token {
	token := &omise.Token{}
	client.MustDo(token, &operations.CreateToken{
		&operations.Card{
			Name:            "Chakrit Wichian",
			Number:          "4242424242424242",
			ExpirationMonth: 2,
			ExpirationYear:  time.Now().AddDate(1, 0, 0).Year(),

			SecurityCode: "123",
			City:         "Bangkok",
			PostalCode:   "10240",
		},
	})

	return token
}

func createTestCharge(client *testutil.TestClient, token *omise.Token) *omise.Charge {
	charge := &omise.Charge{}
	client.MustDo(charge, &operations.CreateCharge{
		Amount:      819229,
		Currency:    "thb",
		Description: "test chrage.",
		Card:        token.ID,
	})

	return charge
}
