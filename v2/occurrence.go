package omise

import (
	"time"
)

// Occurrence represents Occurrence object.
// See https://www.omise.co/occurrences-api for more information.
type Occurrence struct {
	Base
	Location string `json:"location"`
	Message *string `json:"message"`
	ProcessedAt *time.Time `json:"processed_at"`
	Result *string `json:"result"`
	RetryOn Date `json:"retry_on"`
	Schedule string `json:"schedule"`
	ScheduledOn Date `json:"scheduled_on"`
	Status *OccurrenceStatus `json:"status"`
}