package omise

import "time"

// Occurrence represents occurrence charge from Schedule
type Occurrence struct {
	Base
	Schedule     string           `json:"schedule"`
	ScheduleDate time.Time        `json:"schedule_date"`
	RetryDate    time.Time        `json:"retry_date"`
	ProcessedAt  time.Time        `json:"processed_at"`
	Status       OccurrenceStatus `json:"status"`
	Message      *string          `json:"message"`
	Result       string           `json:"result"`
}
