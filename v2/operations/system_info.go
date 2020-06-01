package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	system_info, retrieve := &omise.SystemInfo{}, &RetrieveSystemInfo{
//	}
//	if e := client.Do(system_info, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("system_info: %#v\n", system_info)
//
type RetrieveSystemInfo struct {
}

// Describe
func (req *RetrieveSystemInfo) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/",
		ContentType: "application/json",
	}
}

