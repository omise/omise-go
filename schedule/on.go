package schedule

// WeekDay represents set of weekday
type Weekday string

const (
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
	Saturday  Weekday = "saturday"
	Sunday    Weekday = "sunday"
)

// On represents on field of Schedule object.
type On struct {
	Weekdays       []Weekday
	DaysOfMonth    []int
	WeekdayOfMonth *string
}
