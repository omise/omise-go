package omise

// Weekdays represents an enumeration of possible values for ScheduleOn.
type Weekdays string

// Weekdays can be one of the following list of constants:
const (
	ScheduleOnFriday Weekdays = "friday"
	ScheduleOnMonday Weekdays = "monday"
	ScheduleOnSaturday Weekdays = "saturday"
	ScheduleOnSunday Weekdays = "sunday"
	ScheduleOnThursday Weekdays = "thursday"
	ScheduleOnTuesday Weekdays = "tuesday"
	ScheduleOnWednesday Weekdays = "wednesday"
)

