/*
Package omise provides GO binding for Omise REST API.
Full API documentation is available at https://www.omise.co/docs.
For operations, see the operations subpackage.

Example:

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

*/
package omise
