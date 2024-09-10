package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func GetCapabilities(client *omise.Client) (*omise.Capability, error) {
	capability := &omise.Capability{}

	err := client.Do(capability, &operations.RetrieveCapability{})

	if err != nil {
		return nil, err
	}

	return capability, nil
}
