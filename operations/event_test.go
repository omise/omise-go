package operations_test

import (
	"strings"
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	const (
		EventID    = "evnt_test_526yctupnje5mbldskd"
		TransferID = "trsf_test_526yctqob5djkckq88a"
		EventID2   = "evnt_test_5232t5tlhjwh7nbi14g"
	)

	client := testutil.NewFixedClient(t)

	event := &omise.Event{}
	client.MustDo(event, &RetrieveEvent{EventID})
	r.Equal(t, EventID, event.ID)
	r.Equal(t, "transfer.destroy", event.Key)
	r.Equal(t, TransferID, event.Data.(*omise.Deletion).ID)

	events := &omise.EventList{}
	client.MustDo(events, &ListEvents{})
	r.Len(t, events.Data, 20)

	event = events.Data[0]
	r.Equal(t, EventID2, event.ID)
	r.Equal(t, "customer.create", event.Key)
}

func TestEvent_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	events := &omise.EventList{}
	client.MustDo(events, &ListEvents{})
	r.True(t, len(events.Data) > 1)

	event, retrieve := &omise.Event{}, &RetrieveEvent{events.Data[0].ID}
	client.MustDo(event, retrieve)
	r.True(t, strings.HasPrefix(event.ID, "evnt_test_"))
	r.Equal(t, events.Data[0].ID, event.ID)
}
