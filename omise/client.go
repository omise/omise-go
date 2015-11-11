package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

var client *omise.Client

func getClient() (*omise.Client, error) {
	if client != nil {
		return client, nil
	}

	cl, e := omise.NewClient(config.PKey, config.SKey)
	if e != nil {
		return nil, e
	}

	client = cl
	return client, nil
}

func do(result interface{}, op internal.Operation) error {
	client, e := getClient()
	if e != nil {
		return e
	}

	if e := client.Do(result, op); e != nil {
		return e
	}

	return output(result)
}
