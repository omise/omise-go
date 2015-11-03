# OMISE-GO

Install with:

```go
go get github.com/omise/omise-go
```

# USAGE

Create a client with `omise.NewClient` and use operations object from the
`github.com/omise/omise-go/operations` package to perform API operations.

```go
charge, op := &Charge{}, &operations.CreateCharge{
  Amount: 100000, // ¥10,000
  Currency: "jpy",
  Card: "tok_1234",
}

if e := client.Do(charge, op); e != nil {
  panic(e)
}

fmt.Printf("%#v\n", charge)
```

See documentation on [godoc.org][0] f

# FAQ

### Handling `nil`s

For **requests** to the Omise API: Where optional parameters are concerned, empty values
are considered as not sending the value. Except where doing so may have undesirable
implications.

For **responses** from the Omise API: Optional fields always use nillable types. For
example, an optional `string` field will have type `*string`.

### Test keys

You will need to supply Omise API keys in order to run tests. Refer to the official
documentation on [Authentication][1] for more
information.

To specify the keys, set the environment variables before running the tests as follows:

```go
export OMISE_PUB_KEY=your_public_key
export OMISE_SECRET_KEY=your_secret_key
```

[0]: http://godoc.org/github.com/omise/omise-go
[1]: https://www.omise.co/api-authentication

# LICENSE

See LICENSE file.
