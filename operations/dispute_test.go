package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestDispute(t *testing.T) {
	const (
		DisputeID = "dspt_test_5089off452g5m5te7xs"
	)

	client := testutil.NewFixedClient(t)

	// 4 possible states.
	disputes := &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{})
	if a.Len(t, disputes.Data, 1) {
		a.Equal(t, DisputeID, disputes.Data[0].ID)
	}

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Open})
	if a.Len(t, disputes.Data, 1) {
		a.Equal(t, DisputeID, disputes.Data[0].ID)
		a.Equal(t, omise.Open, disputes.Data[0].Status)
	}

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Pending})
	if a.Len(t, disputes.Data, 1) {
		a.Equal(t, DisputeID, disputes.Data[0].ID)
		a.Equal(t, omise.Pending, disputes.Data[0].Status)
	}

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Closed})
	if a.Len(t, disputes.Data, 1) {
		a.Equal(t, DisputeID, disputes.Data[0].ID)
		a.Equal(t, omise.Won, disputes.Data[0].Status)
	}

	// single instances
	dispute := &omise.Dispute{}
	client.MustDo(dispute, &RetrieveDispute{DisputeID})
	a.Equal(t, DisputeID, dispute.ID)

	client.MustDo(dispute, &UpdateDispute{DisputeID, "Your dispute message"})
	a.Equal(t, "Your dispute message", dispute.Message)
}

func TestDispute_Network(t *testing.T) {
	// TODO: No way to programmatically generates Dispute against the API yet.
	//   so not sure how we can test this thoroughly.
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// only test JSON bindings for now.
	disputes, list := &omise.DisputeList{}, &ListDisputes{}
	client.MustDo(disputes, list)
	if len(disputes.Data) > 0 {
		a.True(t, disputes.Data[0].Status != omise.Pending)
	}
}
