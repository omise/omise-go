/*
Package omise provides GO binding for Omise REST API.
Full REST API documentation is available at https://www.omise.co/docs.

Usage

Create a client with omise.NewClient, supply your public and secret key. Then use the
client.Do method with operation objects from the operations subpackage to perform API
calls. The first parameter to client.Do lets you supply a struct to unmarshal the result
into.

Example:

	client, e := omise.NewClient(OMISE_PUBKEY, OMISE_KEY)
	if e != nil {
		panic(e)
	}

	charge, create := &omise.Charge{}, &operations.CreateCharge{
		Amount:   100000, // ฿1,000.00
		Currency: "thb",
		Card:     "tok_1234",
	}

	if e := client.Do(charge, create); e != nil {
		panic(e)
	}

	fmt.Printf("%#v\n", charge)

Handling nils

For REQUESTS: Where optional parameters are concerned, empty values are considered as not
sending the value except where doing so may have undesirable implications. This is to
avoid the need for pointer indirections when creating the struct.

For RESPONSES: Optional fields always use nillable types or types with well-defined
"empty" value. For example, an optional `string` field will have `*string` type.

Test keys

You will need to supply a valid Omise API keys in order to run tests. Refer to the
official documentation on Authentication at https://www.omise.co/api-authentication for
more information.

To specify the keys, set the environment variables before running the tests as follows:

	export OMISE_PUBKEY=pkey_test_yourpublickey
	export OMISE_KEY=skey_test_yoursecretkey

Live tests

Tests are run against fixtures by default. If you want to run network test against Omise
test servers you can do so by supplying valid test keys *and* setting NETWORK environment
variable to 1.

	cd $GOPATH/src/github.com/omise/omise-go
	NETWORK=1 go test -v ./operations

*/
package omise
