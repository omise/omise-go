package omise

// OccurrenceStatus represents an enumeration of possible values for Occurrence.
type OccurrenceStatus string

// OccurrenceStatus can be one of the following list of constants:
const (
	OccurrenceFailed OccurrenceStatus = "failed"
	OccurrenceScheduled OccurrenceStatus = "scheduled"
	OccurrenceSkipped OccurrenceStatus = "skipped"
	OccurrenceSuccessful OccurrenceStatus = "successful"
)

