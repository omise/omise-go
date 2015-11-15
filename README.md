# OMISE-GO

[![GoDoc](https://godoc.org/github.com/omise/omise-go?status.svg)][0] 
[![Build Status](https://travis-ci.org/omise/omise-go.svg)](https://travis-ci.org/omise/omise-go)

Install with:

```go
go get github.com/omise/omise-go
```

# USAGE

See [godoc.org][0] for usage instruction.

Example:

```go
const (
  OMISE_PUBLIC_KEY = "pkey_test_1234"
  OMISE_SECRET_KEY = "skey_test_1234"
)

client, e := omise.NewClient(OMISE_PUBLIC_KEY, OMISE_SECRET_KEY)
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

# LICENSE

See LICENSE file.

[0]: https://godoc.org/github.com/omise/omise-go
