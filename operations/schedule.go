package operations

import (
	"encoding/json"
	"time"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
	"github.com/omise/omise-go/schedule"
)

// CreateChargeSchedule represent create charge schedule API payload
//
// Example:
//
//	var schd omise.Schedule
//	create := operations.CreateChargeSchedule{
//              Every:  3,
//              Period: schedule.PeriodWeek,
//              Weekdays: []schedule.Weekday{
//              schedule.Monday,
//              	schedule.Saturday,
//              },
//              StartDate: "2017-05-15",
//              EndDate:   "2018-05-15",
//              Customer:  "customer_id",
//              Amount:    100000,
//	}
//	if e := client.Do(&schd, &create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created schedule:", schd.ID)
//
type CreateChargeSchedule struct {
	Every          int
	Period         schedule.Period
	StartDate      string
	EndDate        string
	Weekdays       schedule.Weekdays
	DaysOfMonth    schedule.DaysOfMonth
	WeekdayOfMonth string

	Customer    string
	Amount      int
	Currency    string
	Card        string
	Description string
}

func (req *CreateChargeSchedule) MarshalJSON() ([]byte, error) {
	type charge struct {
		Customer    string `json:"customer"`
		Amount      int    `json:"amount"`
		Currency    string `json:"currency,omitempty"`
		Card        string `json:"card,omitempty"`
		Description string `json:"description,omitempty"`
	}

	type on struct {
		Weekdays       []schedule.Weekday `json:"weekdays,omitempty"`
		DaysOfMonth    []int              `json:"days_of_month,omitempty"`
		WeekdayOfMonth string             `json:"weekday_of_month,omitempty"`
	}

	type param struct {
		Every     int             `json:"every"`
		Period    schedule.Period `json:"period"`
		StartDate *omise.Date     `json:"start_date,omitempty"`
		EndDate   omise.Date      `json:"end_date"`
		On        *on             `json:"on,omitempty"`

		Charge charge `json:"charge"`
	}

	p := param{
		Every:  req.Every,
		Period: req.Period,
		Charge: charge{
			Customer:    req.Customer,
			Amount:      req.Amount,
			Currency:    req.Currency,
			Card:        req.Card,
			Description: req.Description,
		},
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, err
		}
		p.StartDate = (*omise.Date)(&startDate)
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, err
		}
		p.EndDate = omise.Date(endDate)
	}

	switch {
	case p.Period == "week":
		p.On = &on{
			Weekdays: req.Weekdays,
		}
	case p.Period == "month" && req.DaysOfMonth != nil:
		p.On = &on{
			DaysOfMonth: req.DaysOfMonth,
		}
	case p.Period == "month" && req.WeekdayOfMonth != "":
		p.On = &on{
			WeekdayOfMonth: req.WeekdayOfMonth,
		}
	}

	return json.Marshal(p)
}

func (req *CreateChargeSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// CreateTransferSchedule represent create transfer schedule API payload
//
// Example:
//
//	var schd omise.Schedule
//	create = operations.CreateTransferSchedule{
//              Every:  3,
//              Period: schedule.PeriodWeek,
//              Weekdays: []schedule.Weekday{
//              schedule.Monday,
//              	schedule.Saturday,
//              },
//              StartDate: "2017-05-15",
//              EndDate:   "2018-05-15",
//              Recipient:  "recipient_id",
//              Amount:    100000,
//	}
//	if e := client.Do(&schd, &create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created schedule:", schd.ID)
//
type CreateTransferSchedule struct {
	Every          int
	Period         schedule.Period
	StartDate      string
	EndDate        string
	Weekdays       schedule.Weekdays
	DaysOfMonth    schedule.DaysOfMonth
	WeekdayOfMonth string

	Recipient           string
	Amount              int
	PercentageOfBalance float64
}

func (req *CreateTransferSchedule) MarshalJSON() ([]byte, error) {
	type transfer struct {
		Recipient           string  `json:"recipient"`
		Amount              int     `json:"amount,omitempty"`
		PercentageOfBalance float64 `json:"percentage_of_balance,omitempty"`
	}

	type on struct {
		Weekdays       []schedule.Weekday `json:"weekdays,omitempty"`
		DaysOfMonth    []int              `json:"days_of_month,omitempty"`
		WeekdayOfMonth string             `json:"weekday_of_month,omitempty"`
	}

	type param struct {
		Every     int             `json:"every"`
		Period    schedule.Period `json:"period"`
		StartDate *omise.Date     `json:"start_date,omitempty"`
		EndDate   omise.Date      `json:"end_date"`
		On        *on             `json:"on,omitempty"`

		Transfer transfer `json:"transfer"`
	}

	p := param{
		Every:  req.Every,
		Period: req.Period,
		Transfer: transfer{
			Recipient:           req.Recipient,
			Amount:              req.Amount,
			PercentageOfBalance: req.PercentageOfBalance,
		},
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, err
		}
		p.StartDate = (*omise.Date)(&startDate)
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, err
		}
		p.EndDate = omise.Date(endDate)
	}

	switch {
	case p.Period == "week":
		p.On = &on{
			Weekdays: req.Weekdays,
		}
	case p.Period == "month" && req.DaysOfMonth != nil:
		p.On = &on{
			DaysOfMonth: req.DaysOfMonth,
		}
	case p.Period == "month" && req.WeekdayOfMonth != "":
		p.On = &on{
			WeekdayOfMonth: req.WeekdayOfMonth,
		}
	}

	return json.Marshal(p)
}

func (req *CreateTransferSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// ListSchedules represent list schedule API payload
//
// Example:
//
//      var schds omise.ScheduleList
//	list := ListSchedules{
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&schds, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of schedules made in the last hour:", len(schds.Data))
//
type ListSchedules struct {
	List
}

func (req *ListSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules",
		ContentType: "application/json",
	}
}

// ListScheduleOccurrences represent list schedule occurrences API payload
//
// Example:
//
//      var occurrences omise.OccurrenceList
//	list := ListOccurrenceSchedules{
//		ScheduleID: "schd_57z9hj228pusa652nk1",
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&occurrences, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("occurrences made in the last hour:", len(occurrences.Data))
//
type ListScheduleOccurrences struct {
	ScheduleID string
	List
}

func (req *ListScheduleOccurrences) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID + "/occurrences",
		ContentType: "application/json",
	}
}

// ListChargeSchedules represent list charge schedules API payload
//
// Example:
//
//      var schds omise.ScheduleList
//	list = ListChargeSchedules{
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&schds, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of charge schedules made in the last hour:", len(schds.Data))
//
type ListChargeSchedules struct {
	List
}

func (req *ListChargeSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/schedules",
		ContentType: "application/json",
	}
}

// ListTransferSchedules represent list transfer schedules API payload
//
// Example:
//
//      var schds omise.ScheduleList
//	list = ListTransferSchedules{
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&schds, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of transfer schedules made in the last hour:", len(schds.Data))
//
type ListTransferSchedules struct {
	List
}

func (req *ListTransferSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers/schedules",
		ContentType: "application/json",
	}
}

// RetrieveSchedule represent retrieve schedule API payload
//
// Example:
//
//	var schd omise.Schedule
//	if e := client.Do(&schd, &RetrieveSchedule{ScheduleID: "schd_57z9hj228pusa652nk1"}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule #schd_57z9hj228pusa652nk1: %#v\n", schd)
//
type RetrieveSchedule struct {
	ScheduleID string `json:"-"`
}

func (req *RetrieveSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}

// DestroySchedule represent destroy schedule API payload
//
// Example:
//      var deletedSchd omise.Schedule
//	if e := client.Do(&deletedSchd, &DestroySchedule{ScheduleID: "schd_57z9hj228pusa652nk1"}); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("destroyed schedule:", deletedSchd.ID)
//
type DestroySchedule struct {
	ScheduleID string `json:"-"`
}

func (req *DestroySchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/schedules/" + req.ScheduleID,
		ContentType: "application/json",
	}
}
