package operations

import (
	"testing"

	omise "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/require"
)

func TestRetrieveOccurrence(t *testing.T) {
	OccurrenceID := "occu_57z9hj228pusa652nk1"

	client := testutil.NewFixedClient(t)
	var occurrence omise.Occurrence
	client.MustDo(&occurrence, &RetrieveOccurrence{OccurrenceID: OccurrenceID})
	r.Equal(t, OccurrenceID, occurrence.ID)
}

func TestRetrieveOccurrence_Network(t *testing.T) {
	// OccurrenceID must have this occurrence in test server
	OccurrenceID := "occu_57z9hj228pusa652nk1"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	var occurrence omise.Occurrence
	client.MustDo(&occurrence, &RetrieveOccurrence{OccurrenceID: OccurrenceID})

	t.Logf("%#v\n", occurrence)
}
