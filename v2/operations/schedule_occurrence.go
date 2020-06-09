package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Occurrence{}, &ListOccurrences{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListOccurrences struct {
	ScheduleID string `json:"-"`
}

func (req *ListOccurrences) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID + "/occurrences",
		ContentType: "application/json",
	}
}

