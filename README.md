# Omise Go Library

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)][0]
[![omise-go v1](https://github.com/omise/omise-go/actions/workflows/v1-ci.yml/badge.svg)](https://github.com/omise/omise-go/actions/workflows/v1-ci.yml)

Opn Payments helps merchants of any size accept payments online.
This library offers Go integration to the Opn Payments API.

Install with:

```sh
go get github.com/omise/omise-go
```

## Security Warning

**Please do NOT use Omise Go library versions less than 1.0.5, as they are outdated and have security vulnerabilities.**


## Compliance warning

Card data should never transit through your server. This library provides the means to create
tokens on the server side but should only be used for testing or **only if you currently
have a valid PCI-DSS Attestation of Compliance (AoC) delivered by a certified QSA
Auditor.**

Instead, we recommend that you follow our guide on how to safely
[collect credit information](https://docs.opn.ooo/collecting-card-information).

## Version Guidance

This library offers maintenance support for 1.x and active development for 2.x. Please refer to the table for specific version compatibility and requirements.

| Version | Support                 | Omise API Version           | Go Version |
| --------| ----------------------- | --------------------------- | -----------| 
| 1.x     | Maintenance Mode  | <= 2017-11-02               | >= 1.16 |
| 2.x     | Active                  | 2019-05-29                  | >= 1.16 |

<sup><b>*Maintenance mode means bug fixes and security patches only</b></sup>

## Compatibility

Version `v1.2.0` and higher of this package is designed to work with Go version 1.16 or higher. It is not compatible with Go versions 1.15 and lower.

If you are using an older version of Go, please consider upgrading to a compatible version to use this library effectively.

## Usage

For usage instructions, see [godoc.org][0] in tandem with the [Opn Payments API Documentation][1].

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

## API version

You can choose the API version to use with Opn Payments. Each new API version has new features
that might be incompatible with previous versions. You can change the default version by
visiting your Opn Payments Dashboard.

The version configured here will have higher priority than the version set in your Opn Payments
account. This is useful if you have multiple environments with different API versions for
testing. (e.g., Development on the latest version, but production is on an older version).

```
client.APIVersion = "2015-11-06"
```

It is highly recommended to set this version to the current version that you are using. You can
learn more about this feature in our [versioning guide](https://docs.opn.ooo/api-versioning).

## Enable debug mode

You can monitor the response code and body from each API call by utilizing the library's debug mode, which you can turn on or off. To activate this feature, invoke `SetDebug()` immediately after initializing your client.

```
client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
if e != nil {
	log.Fatal(e)
}

// Enabling debug mode to monitor response from API call
client.SetDebug(true)
```

## License

See [LICENSE][2] file.

[0]: https://godoc.org/github.com/omise/omise-go
[1]: https://docs.opn.ooo
[2]: https://raw.githubusercontent.com/omise/omise-go/master/LICENSE
