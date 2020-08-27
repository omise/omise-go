package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Schedule represents Schedule object.
// See https://www.omise.co/schedules-api for more information.
type Schedule struct {
	Base
	Active bool `json:"active"`
	Charge *ChargeSchedule `json:"charge"`
	EndOn Date `json:"end_on"`
	EndedAt time.Time `json:"ended_at"`
	Every int `json:"every"`
	InWords string `json:"in_words"`
	Location string `json:"location"`
	NextOccurrencesOn []string `json:"next_occurrences_on"`
	Occurrences *OccurrenceList `json:"occurrences"`
	On *ScheduleOn `json:"on"`
	Period SchedulePeriod `json:"period"`
	StartOn Date `json:"start_on"`
	Status ScheduleStatus `json:"status"`
	Transfer *TransferSchedule `json:"transfer"`
}

// ScheduleService represents resource service.
type ScheduleService struct {
	*Client
}

// Schedule defines resource service.
func (c *Client) Schedule() *ScheduleService {
	return &ScheduleService{c}
}

// Destroy destroys schedule
//
// Example:
//
//	schedule, err := client.Schedule().Destroy(ctx, &DestroyScheduleParams{
//		ScheduleID: "schd_4hr8b3gd9c"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *ScheduleService) Destroy(ctx context.Context, params *DestroyScheduleParams) (*Schedule, error) {
	result := &Schedule{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyScheduleParams params object.
type DestroyScheduleParams struct {
	ScheduleID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyScheduleParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves schedule
//
// Example:
//
//	schedule, err := client.Schedule().Retrieve(ctx, &RetrieveScheduleParams{
//		ScheduleID: "schd_4hr8b3gd9c"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *ScheduleService) Retrieve(ctx context.Context, params *RetrieveScheduleParams) (*Schedule, error) {
	result := &Schedule{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveScheduleParams params object.
type RetrieveScheduleParams struct {
	ScheduleID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveScheduleParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// List lists schedules
//
// Example:
//
//	schedule, err := client.Schedule().List(ctx, &ListSchedulesParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *ScheduleService) List(ctx context.Context, params *ListSchedulesParams) (*ScheduleList, error) {
	result := &ScheduleList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListSchedulesParams params object.
type ListSchedulesParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListSchedulesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// Create creates schedule
//
// Example:
//
//	schedule, err := client.Schedule().Create(ctx, &CreateScheduleParams{
//		EndDate: nil
//		Every: nil
//		Period: nil
//		StartDate: nil
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *ScheduleService) Create(ctx context.Context, params *CreateScheduleParams) (*Schedule, error) {
	result := &Schedule{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateScheduleParams params object.
type CreateScheduleParams struct {
	Charge *ChargeScheduleParams `json:"charge,omitempty"`
	EndDate DateString `json:"end_date,omitempty"`
	Every int `json:"every"`
	On *ScheduleOnParams `json:"on,omitempty"`
	Period SchedulePeriod `json:"period"`
	StartDate DateString `json:"start_date,omitempty"`
	Transfer *TransferScheduleParams `json:"transfer,omitempty"`
}

// ChargeScheduleParams params object.
type ChargeScheduleParams struct {
	Amount int64 `json:"amount"`
	Card string `json:"card,omitempty"`
	Currency string `json:"currency,omitempty"`
	Customer string `json:"customer"`
	Description string `json:"description,omitempty"`
}

// TransferScheduleParams params object.
type TransferScheduleParams struct {
	Amount int64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	PercentageOfBalance float64 `json:"percentage_of_balance,omitempty"`
	Recipient string `json:"recipient"`
}

// ScheduleOnParams params object.
type ScheduleOnParams struct {
	DaysOfMonth []int `json:"days_of_month,omitempty"`
	WeekdayOfMonth string `json:"weekday_of_month,omitempty"`
	Weekdays []Weekdays `json:"weekdays,omitempty"`
}

// Describe describes structure of request
func (req *CreateScheduleParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// ListOccurrences lists schedule occurrences
//
// Example:
//
//	occurrence, err := client.Schedule().ListOccurrences(ctx, &ListScheduleOccurrencesParams{
//		ScheduleID: "schd_4hr8b3gd9c"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("occurrence: %#v\n", occurrence)
//
func (s *ScheduleService) ListOccurrences(ctx context.Context, params *ListScheduleOccurrencesParams) (*OccurrenceList, error) {
	result := &OccurrenceList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListScheduleOccurrencesParams params object.
type ListScheduleOccurrencesParams struct {
	ScheduleID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListScheduleOccurrencesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID + "/occurrences",
		ContentType: "application/json",
	}
}

