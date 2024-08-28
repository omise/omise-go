package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateCustomer(client *omise.Client) (*omise.Customer, error) {
	customer, createCustomer := &omise.Customer{}, &operations.CreateCustomer{
		Email:       "john.doe@example.com",
		Description: "John Doe (id: 30)",
	}

	if err := client.Do(customer, createCustomer); err != nil {
		return nil, err
	}

	return customer, nil
}
