package src

import (
	"fmt"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateChargeCard(client *omise.Client) (*omise.Charge, error) {
	// Creates a charge from the token
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            "Chakrit Wichian",
		Number:          "4242424242424242",
		ExpirationMonth: 2,
		ExpirationYear:  time.Now().AddDate(1, 0, 0).Year(),
		SecurityCode:    "123",
		Email:           "aashish@omise.co",
		PhoneNumber:     "1234567890",
		Country:         "TH",
		City:            "Bangkok",
		PostalCode:      "10240",
		Street1:         "Sukumvit 63, Street 1",
		Street2:         "Sukumvit 63, Street 2",
	}

	if e := client.Do(token, createToken); e != nil {
		return nil, e
	}

	timeString := "2024-09-15 02:30:45"
	theTime, err := time.Parse("2006-01-02 03:04:05", timeString)

	if err != nil {
		fmt.Println("Could not parse time:", err)
	}

	client.WithCustomHeaders(map[string]string{
		"SUB_MERCHANT_ID": "team_60vxq170pzvgq8w7cz5",
	})

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:           300000, // ฿ 1,000.00
		Currency:         "thb",
		ReturnURI:        "http://omise.co",
		Card:             token.ID,
		WebhookEndpoints: []string{"https://webhook.site/e3e511da-f9ab-4b69-b93d-dd868459a760"},

		// Customer
		Description: "This is a card cahrge",
		DontCapture: false,
		Metadata: map[string]interface{}{
			"Hello": "World",
		},
		ExpiresAt:            &theTime,
		Ip:                   "192.168.1.1",
		TransactionIndicator: omise.MIT,
		RecurringReason:      omise.Blank,
		PlatformFee: omise.PlatformFee{
			Fixed:      10,
			Percentage: 2,
		},
	}

	if e := client.Do(charge, createCharge); e != nil {
		return nil, e
	}

	return charge, nil
}
