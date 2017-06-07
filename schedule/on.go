package schedule

// DaysOfMonth represents slice of day of month
type DaysOfMonth []int

// Weekdays represents slice of weekday
type Weekdays []Weekday

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
	Weekdays       Weekdays    `json:"weekdays"`
	DaysOfMonth    DaysOfMonth `json:"days_of_month"`
	WeekdayOfMonth *string     `json:"weekday_of_month"`
}
