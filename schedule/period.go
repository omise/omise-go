package schedule

// Period represents an enumeration of possible period of a Schedule object.
type Period string

// Period can be one of the following list of constants:
const (
	PeriodDay   Period = "day"
	PeriodWeek  Period = "week"
	PeriodMonth Period = "month"
)
