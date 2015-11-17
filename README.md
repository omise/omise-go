# OMISE-GO

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)][0] 
[![Build Status](https://travis-ci.org/omise/omise-go.svg)](https://travis-ci.org/omise/omise-go)

Install with:

```go
go get github.com/omise/omise-go
```

# USAGE

See [godoc.org][0] in tandem with the [Omise API Documentation][1] for usage instruction.

Example:

```go
package main

import (
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

func main() {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
	}

	// Creates a token from a test card.
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            "OMISE-GO Test Card",
		Number:          "4242424242424242",
		ExpirationMonth: 12,
		ExpirationYear:  2018,
	}
	if e := client.Do(token, createToken); e != nil {
		log.Fatal(e)
	}

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   100000, // ฿ 1,000.00
		Currency: "thb",
		Card:     token.ID,
	}
	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}
```

# LICENSE

See LICENSE file.

[0]: https://godoc.org/github.com/omise/omise-go
[1]: https://www.omise.co/docs
