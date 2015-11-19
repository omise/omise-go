/*
!! EXPERIMENTAL !! This package is experimental and will be subject to frequent changes.

Package api provides convenience methods for performing operations against the Omise REST
API.  Full REST API documentation is available at https://www.omise.co/docs.

Usage

Creates an Omise API client with omise.NewClient() and supply it to the api.With() method.
The return value is a struct with fields reflecting the types defined in the operations
package. For example, operations.CreateCharge{} can be accessed via the API object as
apiobj.Charge.Create() method.

Example:

	client, e := omise.NewClient(OMISE_PUBKEY, OMISE_KEY)
	if e != nil {
		panic(e)
	}

	// Creates a new API object.
	apiobj := api.With(client)

	// Access Charges API
	charge, e := apiobj.Charge.Create(&operations.CreateCharge{
		Amount:   100000, // ฿10,000
		Currency: "thb",
		Card:     "tok_1234",
	})
	if e != nil {
		panic(e)
	}

	fmt.Printf("%#v\n", charge)
*/
package api
