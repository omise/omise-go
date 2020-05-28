package omise

type OccurrenceStatus string

const (
	OccurrenceFailed OccurrenceStatus = "failed"
	OccurrenceScheduled OccurrenceStatus = "scheduled"
	OccurrenceSkipped OccurrenceStatus = "skipped"
	OccurrenceSuccessful OccurrenceStatus = "successful"
)
