**⚠️ This is not OMG, the blockchain. Visit
https://omg.network/ for that. ⚠️**

---

# OMISE GO CLIENT

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)][0]
![omise-go v1](https://github.com/omise/omise-go/workflows/omise-go%20v1/badge.svg)

Omise is a payment service provider currently operating in Thailand. Omise provides a set
of clean APIs that helps merchants of any size accept credit cards online.

This library offers GO integration to the Omise API.

Install with:

```go
go get github.com/omise/omise-go
```

# COMPLIANCE WARNING

Card data should never transit through your server. This library provides means to create
tokens on the server side but should only be used for testing or **only if you currently
have a valid PCI-DSS Attestation of Compliance (AoC) delivered by a certified QSA
Auditor.**

Instead we recommend that you follow our guide on how to safely
[collect credit information](https://www.omise.co/collecting-card-information).

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

  /** Retrieve a token from a request
   * A token should be created from a client side by using our client-side libraries
   * https://www.omise.co/libraries#client-side-libraries
   * More information:
   * - https://www.omise.co/collecting-card-information
   * - https://www.omise.co/security-best-practices
   **/
	token := "tokn_xxxxxxxxxxxxx"

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

# API VERSION

You can choose which API version to use with Omise. Each new API version has new features
and might not be compatible with previous versions. You can change the default version by
visiting your Omise Dashboard.

The version configured here will have higher priority than the version set in your Omise
account. This is useful if you have multiple environments with different API versions for
testing. (e.g. Development on the latest version but production is on an older version).

```
client.APIVersion = "2015-11-06"
```

It is highly recommended to set this version to the current version you're using. You can
learn more about this feature in our [versioning
guide](https://www.omise.co/api-versioning).

# LICENSE

See [LICENSE][2] file.

[0]: https://godoc.org/github.com/omise/omise-go
[1]: https://www.omise.co/docs
[2]: https://raw.githubusercontent.com/omise/omise-go/master/LICENSE
