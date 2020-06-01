package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	occurrence, retrieve := &omise.Occurrence{}, &RetrieveOccurrence{
//		OccurrenceID: "occur_dn8b8b6jt3"
//	}
//	if e := client.Do(occurrence, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("occurrence: %#v\n", occurrence)
//
type RetrieveOccurrence struct {
	OccurrenceID string `json:"-"`
}

// Describe
func (req *RetrieveOccurrence) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/occurrences/" + req.OccurrenceID,
		ContentType: "application/json",
	}
}

