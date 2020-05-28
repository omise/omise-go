package omise

type ScheduleStatus string

const (
	ScheduleActive ScheduleStatus = "active"
	ScheduleDeleted ScheduleStatus = "deleted"
	ScheduleExpired ScheduleStatus = "expired"
	ScheduleExpiring ScheduleStatus = "expiring"
	ScheduleRunning ScheduleStatus = "running"
	ScheduleSuspended ScheduleStatus = "suspended"
)
