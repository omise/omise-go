package omise

// SchedulePeriod represents an enumeration of possible period of a Schedule object.
type SchedulePeriod string

// SchedulePeriod can be one of the following list of constants:
const (
	SchedulePeriodDay   SchedulePeriod = "day"
	SchedulePeriodWeek  SchedulePeriod = "week"
	SchedulePeriodMonth SchedulePeriod = "month"
)
