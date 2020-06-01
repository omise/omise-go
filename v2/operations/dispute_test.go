package operations_test

import (
	"encoding/json"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestDispute(t *testing.T) {
	const (
		DisputeID = "dspt_test_5089off452g5m5te7xs"
	)

	client := testutil.NewFixedClient(t)

	// 4 possible states.
	disputes := &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{})
	r.Len(t, disputes.Data, 1)
	r.Equal(t, DisputeID, disputes.Data[0].ID)

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Open})
	r.Len(t, disputes.Data, 1)
	r.Equal(t, DisputeID, disputes.Data[0].ID)

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Pending})
	r.Len(t, disputes.Data, 1)
	r.Equal(t, DisputeID, disputes.Data[0].ID)

	disputes = &omise.DisputeList{}
	client.MustDo(disputes, &ListDisputes{State: omise.Closed})
	r.Len(t, disputes.Data, 1)
	r.Equal(t, DisputeID, disputes.Data[0].ID)

	// single instances
	dispute := &omise.Dispute{}
	client.MustDo(dispute, &RetrieveDispute{DisputeID})
	r.Equal(t, DisputeID, dispute.ID)

	client.MustDo(dispute, &UpdateDispute{
		DisputeID: DisputeID,
		Message:   "Your dispute message",
	})
	r.Equal(t, "Your dispute message", dispute.Message)
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
		r.True(t, disputes.Data[0].Status != omise.Pending)
	}
}

func TestUpdateDisputeMarshal_WithMetadata(t *testing.T) {
	req := &UpdateDispute{
		DisputeID: "dspt_test_5089off452g5m5te7xs",
		Message:   "Your dispute message",
		Metadata:  map[string]interface{}{"color": "red"},
	}

	expected := `{"message":"Your dispute message","metadata":{"color":"red"}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateDisputeMarshal_WithoutMetadata(t *testing.T) {
	req := &UpdateDispute{
		DisputeID: "dspt_test_5089off452g5m5te7xs",
		Message:   "Your dispute message",
	}

	expected := `{"message":"Your dispute message"}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}
