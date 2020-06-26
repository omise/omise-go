package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
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
	Status OccurrenceStatus `json:"status"`
}

// OccurrenceService represents resource service.
type OccurrenceService struct {
	*Client
}

// Occurrence defines resource service.
func (c *Client) Occurrence() *OccurrenceService {
	return &OccurrenceService{c}
}

// Retrieve retrieves occurrence
//
// Example:
//
//	occurrence, retrieve := &omise.Occurrence{}, &RetrieveOccurrence{
//		OccurrenceID: "occur_dn8b8b6jt3"
//	}
//	if e := client.Do(occurrence, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("occurrence: %#v\n", occurrence)
//
func (s *OccurrenceService) Retrieve(ctx context.Context, params *RetrieveOccurrenceParams) (*Occurrence, error) {
	result := &Occurrence{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveOccurrenceParams params object.
type RetrieveOccurrenceParams struct {
	OccurrenceID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveOccurrenceParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/occurrences/" + req.OccurrenceID,
		ContentType: "application/json",
	}
}

