### nil fields in requests and resposnes.

For **requests** to the Omise API: Where optional parameters are concerned, empty values
are considered as not sending the value. Except where doing so may have undesirable
implications.

For **responses** from the Omise API: Optional fields always use nillable types. For
example, an optional `string` field will have type `*string`.

### test keys

You will need to supply Omise API keys in order to run tests. Refer to the official
documentation on [Authentication](https://www.omise.co/api-authentication) for more
information.

To specify the keys, set the environment variables before running the tests as follows:

```go
export OMISE_PUB_KEY=your_public_key
export OMISE_SECRET_KEY=your_secret_key
```

