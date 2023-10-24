# OMISE GO CLIENT

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)][0]
[![omise-go v1](https://github.com/omise/omise-go/actions/workflows/v1-ci.yml/badge.svg)](https://github.com/omise/omise-go/actions/workflows/v1-ci.yml)

Opn Payments helps merchants of any size accept payments online.
This library offers Go integration to the Opn Payments API.

Install with:

```sh
go get github.com/omise/omise-go
```

# COMPLIANCE WARNING

Card data should never transit through your server. This library provides means to create
tokens on the server side but should only be used for testing or **only if you currently
have a valid PCI-DSS Attestation of Compliance (AoC) delivered by a certified QSA
Auditor.**

Instead we recommend that you follow our guide on how to safely
[collect credit information](https://docs.opn.ooo/collecting-card-information).

## Compatibility

The version `1.2.0` is designed to work with Go version 1.16 or higher. It is not compatible with Go versions 1.15 and below.

If you are using an older version of Go, please consider upgrading to a compatible version to use this project effectively.


# USAGE

See [godoc.org][0] in tandem with the [Opn Payments API Documentation][1] for usage instruction.

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
	OmisePublicKey = "pkey_test_no1t4tnemucod0e51mo"
	OmiseSecretKey = "skey_test_no1t4tnemucod0e51mo"
)

func main() {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
	}

  /** Retrieve a token from a request
   * A token should be created from a client side by using our client-side libraries
   * https://docs.opn.ooo/libraries#client-side-libraries
   * More information:
   * - https://docs.opn.ooo/collecting-card-information
   * - https://docs.opn.ooo/security-best-practices
   **/
	token := "tokn_test_no1t4tnemucod0e51mo"

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   100000, // ฿ 1,000.00
		Currency: "thb",
		Card:     token,
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
visiting your Opn Payments Dashboard.

The version configured here will have higher priority than the version set in your Omise
account. This is useful if you have multiple environments with different API versions for
testing. (e.g. Development on the latest version but production is on an older version).

```
client.APIVersion = "2015-11-06"
```

It is highly recommended to set this version to the current version you're using. You can
learn more about this feature in our [versioning guide](https://docs.opn.ooo/api-versioning).

# LICENSE

See [LICENSE][2] file.

[0]: https://godoc.org/github.com/omise/omise-go
[1]: https://docs.opn.ooo
[2]: https://raw.githubusercontent.com/omise/omise-go/master/LICENSE
