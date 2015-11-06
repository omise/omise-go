# OMISE-GO

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)](https://godoc.org/github.com/omise/omise-go)

Install with:

```go
go get github.com/omise/omise-go
```

# USAGE

Create a client with `omise.NewClient` and make operations object from the
`github.com/omise/omise-go/operations` package to perform API operations.

```go
client, e := omise.NewClient(OMISE_PUBKEY, OMISE_KEY)
if e != nil {
  panic(e)
}

charge, create := &omise.Charge{}, &operations.CreateCharge{
  Amount: 100000, // ¥10,000
  Currency: "jpy",
  Card: "tok_1234",
}

if e := client.Do(charge, create); e != nil {
  panic(e)
}

fmt.Printf("%#v\n", charge)
```

The `client.Do` method accepts a struct to unmarshal the response into and the operation
to perform.  See documentation on [godoc.org][0] for the full API documentation.

# FAQ

### Handling `nil`s

For **REQUESTS**: Where optional parameters are concerned, empty values
are considered as not sending the value. Except where doing so may have undesirable
implications.

For **RESPONSES**: Optional fields always use nillable types. For
example, an optional `string` field will have type `*string`.

### Test keys

You will need to supply Omise API keys in order to run network tests, which will cause
changes to your Omise test account. Refer to the official documentation on
[Authentication][1] for more information.

To specify the keys, set the environment variables before running the tests as follows:

```go
export OMISE_PUB_KEY=your_public_key
export OMISE_SECRET_KEY=your_secret_key
```

### Live tests

Tests are run against fixtures by default, if you want to test this against Omise test
servers you can do so by supplying valid test keys *an* setting `export NETWORK=1` before
running `go test` or `make test`

[0]: http://godoc.org/github.com/omise/omise-go
[1]: https://www.omise.co/api-authentication

# LICENSE

See LICENSE file.
