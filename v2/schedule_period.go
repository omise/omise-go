package omise

// SchedulePeriod represents an enumeration of possible values for Schedule.
type SchedulePeriod string

// SchedulePeriod can be one of the following list of constants:
const (
	ScheduleDay SchedulePeriod = "day"
	ScheduleMonth SchedulePeriod = "month"
	ScheduleWeek SchedulePeriod = "week"
)

