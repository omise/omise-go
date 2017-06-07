package omise

import "github.com/omise/omise-go/schedule"

// Schedule represents Omise's schedule object.
// See https://www.omise.co/schedule-api for more information.
type Schedule struct {
	Base
	Status          schedule.Status          `json:"status"`
	Every           int                      `json:"every"`
	Period          schedule.Period          `json:"period"`
	On              schedule.On              `json:"on"`
	InWords         string                   `json:"in_words"`
	StartDate       Date                     `json:"start_date"`
	EndDate         Date                     `json:"end_date"`
	Charge          *schedule.ChargeDetail   `json:"charge"`
	Transfer        *schedule.TransferDetail `json:"transfer"`
	Occurrences     OccurrenceList           `json:"occurrences"`
	NextOccurrences []Date                   `json:"next_occurrences"`
}
