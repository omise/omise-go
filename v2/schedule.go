package omise

import (
	"time"
)

// Schedule represents Schedule object.
// See https://www.omise.co/schedules-api for more information.
type Schedule struct {
	Base
	Active bool `json:"active"`
	Charge *ChargeScheduling `json:"charge"`
	EndOn Date `json:"end_on"`
	EndedAt time.Time `json:"ended_at"`
	Every int `json:"every"`
	InWords string `json:"in_words"`
	Location string `json:"location"`
	NextOccurrencesOn []string `json:"next_occurrences_on"`
	Occurrences *OccurrenceList `json:"occurrences"`
	On *ScheduleOn `json:"on"`
	Period *SchedulePeriod `json:"period"`
	StartOn Date `json:"start_on"`
	Status *ScheduleStatus `json:"status"`
	Transfer *TransferScheduling `json:"transfer"`
}
