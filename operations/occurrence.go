package operations

import "github.com/omise/omise-go/internal"

// RetrieveOccurrence represent retrieve occurrence API payload
//
// Example:
//
//	var occurrence omise.Occurrence
//	if e := client.Do(&occurrence, &RetrieveOccurrence{OccurrenceID: "occu_57z9hj228pusa652nk1"}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("occurrence #occu_57z9hj228pusa652nk1: %#v\n", occurrence)
//
type RetrieveOccurrence struct {
	OccurrenceID string `json:"-"`
}

func (req *RetrieveOccurrence) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/occurrences/" + req.OccurrenceID,
		ContentType: "application/json",
	}
}
