package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.SystemInfo{}, &Show{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type Show struct {
}

func (req *Show) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        ,
		ContentType: "application/json",
	}
}

