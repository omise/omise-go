package schedule

// OccurrenceStatus represents an enumeration of possible status of a Occurrence object.
type OccurrenceStatus string

// OccurrenceStatus can be one of the following list of constants:
const (
	OccurrenceSkip       OccurrenceStatus = "skipped"
	OccurrenceFailed     OccurrenceStatus = "failed"
	OccurrenceSuccessful OccurrenceStatus = "successful"
)
