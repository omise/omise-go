package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	client := testutil.NewFixedClient(t)
	capability := &omise.Capability{}
	client.MustDo(capability, &RetrieveCapability{})
	r.Equal(t, capability.Country, "TH")
}

func TestCapability_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	capability := &omise.Capability{}
	client.MustDo(capability, &RetrieveCapability{})
	r.Equal(t, capability.Object, "capability")

	testutil.LogObj(t, capability)
}
