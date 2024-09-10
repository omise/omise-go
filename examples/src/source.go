package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateSource(client *omise.Client) (*omise.Source, error) {
	result := &omise.Source{}

	err := client.Do(result, &operations.CreateSource{
		Type:     "promptpay",
		Amount:   100000,
		Currency: "THB",
		Billing: omise.BillingShipping{
			Country:    "TH",
			City:       "Bangkok",
			PostalCode: "10260",
			State:      "Bangkok",
			Street1:    "Sukumvit 63",
			Street2:    "Sukumvit 63",
		},
		Shipping: omise.BillingShipping{
			Country:    "TH",
			City:       "Chang Mai",
			PostalCode: "10260",
			State:      "Chang Mai",
			Street1:    "Sukumvit 63",
			Street2:    "Sukumvit 63",
		},
		PromotionCode: "Hello",
		Items: []omise.Items{
			{
				Amount:   100000,
				Sku:      "sku-110",
				Name:     "Book",
				Quantity: 1,
				Category: "Reding",
				Brand:    "No brand",
				ItemUri:  "https://assets2.staging-omise.co/assets/external-logo-0059ea57dfd850a136e2fa75871113674eed0a23ca5340bded0d994fd674e9f2.svg",
				ImageUri: "https://assets2.staging-omise.co/assets/external-logo-0059ea57dfd850a136e2fa75871113674eed0a23ca5340bded0d994fd674e9f2.svg",
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
