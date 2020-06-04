package operations


// Example:
//
//	charge, update := &omise.Occurrence{}, &RetrieveOccurrence{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveOccurrence struct {
	Base
	OccurrenceID string `json:"-"`
}

func (req *RetrieveOccurrence) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/occurrences/" + req.OccurrenceID,
		ContentType: "application/json",
	}
}

