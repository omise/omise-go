package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	capability, retrieve := &omise.Capability{}, &RetrieveCapability{
//	}
//	if e := client.Do(capability, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("capability: %#v\n", capability)
//
type RetrieveCapability struct {
}

// Describe
func (req *RetrieveCapability) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/capability",
		ContentType: "application/json",
	}
}

