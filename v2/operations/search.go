package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Search{}, &SearchSearch{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchSearch struct {
}

func (req *SearchSearch) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/search",
		ContentType: "application/json",
	}
}

