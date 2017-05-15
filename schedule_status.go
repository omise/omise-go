package omise

// ScheduleStatus represents an enumeration of possible status of a Schedule object.
type ScheduleStatus string

// ScheduleStatus can be one of the following list of constants:
const (
	ScheduleActive    ScheduleStatus = "active"
	ScheduleExpiring  ScheduleStatus = "expiring"
	ScheduleExpired   ScheduleStatus = "expired"
	ScheduleDeleted   ScheduleStatus = "deleted"
	ScheduleSuspended ScheduleStatus = "suspended"
)
