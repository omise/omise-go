package omise

// ScheduleOn represents resource object.
type ScheduleOn struct {
	Base
	DaysOfMonth []int `json:"days_of_month"`
	WeekdayOfMonth string `json:"weekday_of_month"`
	Weekdays []Weekdays `json:"weekdays"`
}

