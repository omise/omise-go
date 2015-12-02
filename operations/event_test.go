package operations_test

import (
	"strings"
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
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
	a.Equal(t, EventID, event.ID)
	a.Equal(t, "transfer.destroy", event.Key)
	a.Equal(t, TransferID, event.Data.(*omise.Deletion).ID)

	events := &omise.EventList{}
	client.MustDo(events, &ListEvents{})
	a.Len(t, events.Data, 20)

	event = events.Data[0]
	a.Equal(t, EventID2, event.ID)
	a.Equal(t, "customer.create", event.Key)
}

func TestEvent_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	events := &omise.EventList{}
	client.MustDo(events, &ListEvents{})
	a.True(t, len(events.Data) > 1)

	event, retrieve := &omise.Event{}, &RetrieveEvent{events.Data[0].ID}
	client.MustDo(event, retrieve)
	a.True(t, strings.HasPrefix(event.ID, "evnt_test_"))
	a.Equal(t, events.Data[0].ID, event.ID)
}
