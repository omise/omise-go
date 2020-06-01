package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

// Example:
//
//	schedule, destroy := &omise.Schedule{}, &DestroySchedule{
//		ScheduleID: "schd_4hr8b3gd9c"
//	}
//	if e := client.Do(schedule, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type DestroySchedule struct {
	ScheduleID string `json:"-"`
}

// Describe
func (req *DestroySchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// Example:
//
//	schedule, retrieve := &omise.Schedule{}, &RetrieveSchedule{
//		ScheduleID: "schd_4hr8b3gd9c"
//	}
//	if e := client.Do(schedule, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type RetrieveSchedule struct {
	ScheduleID string `json:"-"`
}

// Describe
func (req *RetrieveSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// Example:
//
//	schedule, list := &omise.Schedule{}, &ListSchedules{
//	}
//	if e := client.Do(schedule, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type ListSchedules struct {
	List
}

// Describe
func (req *ListSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	schedule, create := &omise.Schedule{}, &CreateSchedule{
//		EndDate: nil
//		Every: nil
//		Period: nil
//		StartDate: nil
//	}
//	if e := client.Do(schedule, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type CreateSchedule struct {
	Charge *ChargeScheduling `json:"charge,omitempty"`
	EndDate omise.DateString `json:"end_date,omitempty"`
	Every int `json:"every"`
	On *ScheduleOn `json:"on,omitempty"`
	Period omise.SchedulePeriod `json:"period"`
	StartDate omise.DateString `json:"start_date,omitempty"`
	Transfer *TransferScheduling `json:"transfer,omitempty"`
}

// ChargeScheduling
type ChargeScheduling struct {
	Amount int64 `json:"amount"`
	Card string `json:"card,omitempty"`
	Currency string `json:"currency,omitempty"`
	Customer string `json:"customer"`
	Description string `json:"description,omitempty"`
}

// TransferScheduling
type TransferScheduling struct {
	Amount int64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	PercentageOfBalance float64 `json:"percentage_of_balance,omitempty"`
	Recipient string `json:"recipient"`
}

// ScheduleOn
type ScheduleOn struct {
	DaysOfMonth []int `json:"days_of_month,omitempty"`
	WeekdayOfMonth string `json:"weekday_of_month,omitempty"`
	Weekdays []omise.Weekdays `json:"weekdays,omitempty"`
}

// Describe
func (req *CreateSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	occurrence, listOccurrences := &omise.Occurrence{}, &ListScheduleOccurrences{
//		ScheduleID: "schd_4hr8b3gd9c"
//	}
//	if e := client.Do(occurrence, listOccurrences); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("occurrence: %#v\n", occurrence)
//
type ListScheduleOccurrences struct {
	ScheduleID string `json:"-"`
	List
}

// Describe
func (req *ListScheduleOccurrences) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID + "/occurrences",
		ContentType: "application/json",
	}
}

