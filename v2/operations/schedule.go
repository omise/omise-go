package operations


// Example:
//
//	charge, update := &omise.Schedule{}, &ListSchedules{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListSchedules struct {
	Base
}

func (req *ListSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &DestroySchedule{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroySchedule struct {
	Base
	ScheduleID string `json:"-"`
}

func (req *DestroySchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &RetrieveSchedule{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveSchedule struct {
	Base
	ScheduleID string `json:"-"`
}

func (req *RetrieveSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &CreateSchedule{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateSchedule struct {
	Base
	Charge *Charge `json:"charge"`
	EndDate Date `json:"end_date"`
	Every int `json:"every"`
	On *On `json:"on"`
	Period *Period `json:"period"`
	StartDate Date `json:"start_date"`
	Transfer *Transfer `json:"transfer"`
}

func (req *CreateSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Occurrence{}, &ListSchedulesOccurrences{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListSchedulesOccurrences struct {
	Base
	ScheduleID string `json:"-"`
}

func (req *ListSchedulesOccurrences) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/schedules/" + req.ScheduleID + "/occurrences",
		ContentType: "application/json",
	}
}

