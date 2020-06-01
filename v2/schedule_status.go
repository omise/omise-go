package omise

// ScheduleStatus represents an enumeration of possible values for Schedule.
type ScheduleStatus string

// ScheduleStatus can be one of the following list of constants:
const (
	ScheduleActive ScheduleStatus = "active"
	ScheduleDeleted ScheduleStatus = "deleted"
	ScheduleExpired ScheduleStatus = "expired"
	ScheduleExpiring ScheduleStatus = "expiring"
	ScheduleRunning ScheduleStatus = "running"
	ScheduleSuspended ScheduleStatus = "suspended"
)

