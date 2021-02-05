package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	capability := &omise.Capability{}
//	if e := client.Do(capability, &RetrieveCapability{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("my capability!: %#v\n", capability)
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
