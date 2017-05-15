package omise

// ScheduleOn represents on field of Schedule object.
type ScheduleOn struct {
	Weekdays       []int   `json:"weekdays"`
	DaysOfMonth    []int   `json:"day_of_month"`
	WeekdayOfMonth *string `json:"weekday_of_month"`
}
