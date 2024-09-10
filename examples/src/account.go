package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func GetAccount(client *omise.Client) (*omise.Account, error) {
	account := &omise.Account{}

	err := client.Do(account, &operations.RetrieveAccount{})

	if err != nil {
		return nil, err
	}

	return account, nil
}
