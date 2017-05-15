package omise

import "time"

// Schedule represents Omise's schedule object.
// See https://www.omise.co/schedule-api for more information.
type Schedule struct {
	Base
	Status      ScheduleStatus        `json:"status"`
	Every       int                   `json:"every"`
	Period      SchedulePeriod        `json:"period"`
	On          ScheduleOn            `json:"on"`
	InWords     string                `json:"in_words"`
	StartDate   time.Time             `json:"start_date"`
	EndDate     time.Time             `json:"end_date"`
	Charge      *ChargeScheduleDetail `json:"charge"`
	Occurrences *OccurrenceList       `json:"occurrences"`
}
