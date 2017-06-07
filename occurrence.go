package omise

import (
	"time"

	"github.com/omise/omise-go/schedule"
)

// Occurrence represents occurrence charge from Schedule
type Occurrence struct {
	Base
	Schedule     string                    `json:"schedule"`
	ScheduleDate Date                      `json:"schedule_date"`
	RetryDate    Date                      `json:"retry_date"`
	ProcessedAt  time.Time                 `json:"processed_at"`
	Status       schedule.OccurrenceStatus `json:"status"`
	Message      string                    `json:"message"`
	Result       string                    `json:"result"`
}
