package omise

type ScheduleOn struct {
	Base
	DaysOfMonth []interface{} `json:"days_of_month"`
	WeekdayOfMonth string `json:"weekday_of_month"`
	Weekdays []*Weekdays `json:"weekdays"`
}
