package operations

import "github.com/omise/omise-go/internal"

// Example: Retrieve Capability
//
//	capability, retrieve := &omise.RetrieveCapability{}, &RetrieveCapability{}
//	if e := client.Do(capability, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("list payment methods: %#v\n", capability.PaymentMethods)
//
type RetrieveCapability struct{}

func (req *RetrieveCapability) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/capability",
		ContentType: "application/json",
	}
}
