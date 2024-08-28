package src

import (
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateChargeQR(client *omise.Client) (*omise.Charge, error) {
	source, err := CreateSource(client)

	if err != nil {
		log.Fatal(err)
	}

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:           100000, // ฿ 1,000.00
		Currency:         "thb",
		ReturnURI:        "http://opn.ooo",
		WebhookEndpoints: []string{"https://webhook.site/e3e511da-f9ab-4b69-b93d-dd868459a760"},
		Source:           source.ID,
	}

	if e := client.Do(charge, createCharge); e != nil {
		return nil, e
	}

	return charge, nil
}
