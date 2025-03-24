package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateRecipient(client *omise.Client) (*omise.Recipient, error) {
	recipient, createRecipient := &omise.Recipient{}, &operations.CreateRecipient{
		Name:        "Omise",
		Email:       "aashish@omise.co",
		Description: "This is a test recipient",
		Type:        omise.Corporation,
		TaxID:       "tax_123",
		BankAccount: &omise.BankAccountRequest{
			Brand:    "bbl",
			Number:   "acc12345",
			Name:     "Aashish",
			BankCode: "bbl",
			Type:     omise.Current,
		},
		Metadata: map[string]interface{}{
			"Hello": "World",
		},
	}

	if e := client.Do(recipient, createRecipient); e != nil {
		return nil, e
	}

	return recipient, nil
}
