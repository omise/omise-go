package operations_test

import (
	"context"
	"strings"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	const (
		EventID    = "evnt_test_526yctupnje5mbldskd"
		TransferID = "trsf_test_526yctqob5djkckq88a"
		EventID2   = "evnt_test_5232t5tlhjwh7nbi14g"
	)

	client := testutil.NewFixedClient(t)

	event, _ := client.Event().Retrieve(context.Background(), &omise.RetrieveEventParams{EventID})
	r.Equal(t, EventID, event.ID)
	r.Equal(t, "transfer.destroy", event.Key)
	r.Equal(t, TransferID, event.Data.(*omise.Deletion).ID)

	events := &omise.EventList{}
	events, _ = client.Event().List(context.Background(), &omise.ListEventsParams{})
	r.Len(t, events.Data, 20)

	event = events.Data[0]
	r.Equal(t, EventID2, event.ID)
	r.Equal(t, "customer.create", event.Key)
}

func TestEvent_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	events := &omise.EventList{}
	events, _ = client.Event().List(context.Background(), &omise.ListEventsParams{})
	r.True(t, len(events.Data) > 1)

	event, retrieve := &omise.Event{}, &RetrieveEvent{events.Data[0].ID}
	client.MustDo(event, retrieve)
	r.True(t, strings.HasPrefix(event.ID, "evnt_test_"))
	r.Equal(t, events.Data[0].ID, event.ID)
}
